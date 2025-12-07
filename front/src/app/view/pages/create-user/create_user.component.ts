import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-create-user',
  standalone: true,
  imports: [CommonModule, FormsModule],
  template: `
    <section class="create-user">
      <h1>Créer un nouvel utilisateur</h1>
      <form (ngSubmit)="createUser()">
        <label>
          Prénom:
          <input type="text" [(ngModel)]="firstName" name="firstName" required />
        </label>

        <label>
          Nom:
          <input type="text" [(ngModel)]="lastName" name="lastName" required />
        </label>

        <label>
          Email:
          <input type="email" [(ngModel)]="email" name="email" required />
        </label>

        <label>
          Mot de passe:
          <input type="password" [(ngModel)]="password" name="password" required />
        </label>

        <label>
          Rôle:
          <select [(ngModel)]="role" name="role" required>
            <option value="EMPLOYEE">EMPLOYEE</option>
            <option value="ADMIN">ADMIN</option>
            <option value="MANAGER">MANAGER</option>
          </select>
        </label>

        <label>
          Équipe:
          <select [(ngModel)]="teamId" name="teamId" required>
            <option *ngFor="let team of teams" [value]="team._id">{{ team.name }}</option>
          </select>
        </label>

        <button type="submit">Créer</button>
        <p *ngIf="message" class="message">{{ message }}</p>
      </form>
    </section>
  `,
  styleUrls: ['./create_user.component.css']
})
export class CreateUserComponent {
  firstName = '';
  lastName = '';
  email = '';
  password = '';
  role = 'EMPLOYEE';
  teamId = '';
  teams: any[] = [];
  message: string | null = null;

  constructor(private apiService: ApiService, private router: Router) {
    // Récupérer toutes les équipes pour le select
    this.apiService.getAllTeams().subscribe({
      next: (data) => this.teams = data,
      error: (err) => console.error("Erreur récupération équipes", err)
    });
  }

  createUser() {
    const userData = {
      FirstName: this.firstName,
      LastName: this.lastName,
      Email: this.email,
      Password: this.password,
      Role: this.role,
      Team: this.teamId
    };

    this.apiService.createUser(userData).subscribe({
      next: () => {
        this.message = " Utilisateur créé avec succès !";
        setTimeout(() => this.router.navigate(['/home']), 1500);
      },
      error: (err) => {
        console.error(err);
        this.message = " Erreur lors de la création de l'utilisateur.";
      }
    });
  }
}
