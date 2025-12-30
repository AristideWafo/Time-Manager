import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-create-user',
  standalone: true,
  imports: [CommonModule, FormsModule],
  template: `
    <section class="create-user">
      <h1>Créer un nouvel utilisateur</h1>

      <form (ngSubmit)="createUser()" #userForm="ngForm">

        <label>
          Prénom :
          <input
            type="text"
            [(ngModel)]="firstName"
            name="firstName"
            required
          />
        </label>

        <label>
          Nom :
          <input
            type="text"
            [(ngModel)]="lastName"
            name="lastName"
            required
          />
        </label>

        <label>
          Email :
          <input
            type="email"
            [(ngModel)]="email"
            name="email"
            required
          />
        </label>

        <label>
          Mot de passe :
          <input
            type="password"
            [(ngModel)]="password"
            name="password"
            required
          />
        </label>

        <label>
          Rôle :
          <select [(ngModel)]="role" name="role" required>
            <option value="EMPLOYEE">EMPLOYEE</option>
            <option value="MANAGER">MANAGER</option>
            <option value="ADMIN">ADMIN</option>
          </select>
        </label>

        <label>
          Équipe :
          <select [(ngModel)]="teamId" name="teamId" required>
            <option value="" disabled>-- Sélectionner une équipe --</option>
            <option *ngFor="let team of teams" [value]="team._id">
              {{ team.Name }}
            </option>
          </select>
        </label>

        <button type="submit" [disabled]="userForm.invalid">
          Créer
        </button>

        <p *ngIf="message" class="message">{{ message }}</p>

      </form>
    </section>
  `,
  styleUrls: ['./create_user.component.css']
})
export class CreateUserComponent implements OnInit {

  // Champs formulaire
  firstName = '';
  lastName = '';
  email = '';
  password = '';
  role: 'EMPLOYEE' | 'MANAGER' | 'ADMIN' = 'EMPLOYEE';

  // Team
  teamId = '';
  teams: Array<{ _id: string; Name: string }> = [];

  // UI
  message: string | null = null;

  constructor(
    private apiService: ApiService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.loadTeams();
  }

  // 🔹 Récupération des teams existantes
  loadTeams(): void {
    this.apiService.getAllTeams().subscribe({
      next: (data: any) => {
        console.log('Teams récupérées :', data);
        this.teams = data.teams; // ✅ important
      },
      error: (err) => {
        console.error('Erreur récupération équipes', err);
      }
    });
  }

  // 🔹 Création user + assignation team
  createUser(): void {

    const userData = {
      FirstName: this.firstName,
      LastName: this.lastName,
      Email: this.email,
      Password: this.password,
      Role: this.role
    };

    // 1️⃣ Création de l'utilisateur
    this.apiService.createUser(userData).subscribe({
      next: (createdUser: any) => {
        console.log('User créé :', createdUser);

        // ✅ récupérer l'_id réel dans le backend
        const userId = createdUser.user?._id;

        if (!userId) {
          this.message = "Erreur : ID utilisateur non reçu du backend";
          return;
        }

        // 2️⃣ Assignation de la team
        this.apiService.updateUserTeam(userId, { teamId: this.teamId }).subscribe({
          next: () => {
            this.message = 'Utilisateur créé et assigné à une équipe ✅';
            setTimeout(() => this.router.navigate(['/home']), 1500);
          },
          error: (err) => {
            console.error('Erreur assignation team', err);
            this.message =
              'Utilisateur créé mais erreur lors de l’assignation de la team ❌';
          }
        });
      },
      error: (err) => {
        console.error('Erreur création utilisateur', err);
        this.message = 'Erreur lors de la création de l’utilisateur ❌';
      }
    });
  }
}
