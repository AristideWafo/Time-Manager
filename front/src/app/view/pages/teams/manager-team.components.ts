import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-manager-team',
  standalone: true,
  imports: [CommonModule],
  template: `
    <section>
      <h1>Utilisateurs de votre équipe</h1>

      <div *ngIf="loading">Chargement...</div>
      <div *ngIf="error" class="error">{{ error }}</div>

      <table *ngIf="users.length > 0">
        <thead>
          <tr>
            <th>Prénom</th>
            <th>Nom</th>
            <th>Email</th>
            <th>Rôle</th>
          </tr>
        </thead>
        <tbody>
          <tr *ngFor="let user of users">
            <td>{{ user.FirstName }}</td>
            <td>{{ user.LastName }}</td>
            <td>{{ user.Email }}</td>
            <td>{{ user.Role }}</td>
          </tr>
        </tbody>
      </table>

      <div *ngIf="users.length === 0 && !loading">Aucun utilisateur trouvé.</div>
    </section>
  `
})
export class ManagerTeamComponent implements OnInit {
  users: any[] = [];
  loading = false;
  error: string | null = null;

  constructor(private api: ApiService) { }

  ngOnInit(): void {
    this.loadTeamUsers();
  }

  loadTeamUsers() {
    this.loading = true;
    this.api.getManagerTeamUsers().subscribe({
      next: (res) => {
        this.users = res;
        this.loading = false;
      },
      error: (err) => {
        console.error('Erreur API manager/team/users :', err);
        this.error = 'Impossible de charger les utilisateurs de votre équipe';
        this.loading = false;
      }
    });
  }
}
