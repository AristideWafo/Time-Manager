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
    console.log('🟢 ManagerTeamComponent init');

    const teamId = localStorage.getItem('Team');
    console.log('ℹ️ TeamId récupéré du localStorage :', teamId);

    if (!teamId) {
      this.error = '❌ Aucun teamId trouvé dans le localStorage.';
      return;
    }

    this.loadTeamUsers(teamId);
  }

  loadTeamUsers(teamId: string) {
    this.loading = true;

    // On récupère toutes les équipes
    this.api.getAllTeams().subscribe({
      next: (res: any) => {
        console.log('ℹ️ Teams récupérées :', res);

        if (!res.teams || !Array.isArray(res.teams)) {
          this.error = '❌ Aucune équipe trouvée dans la réponse API.';
          this.loading = false;
          return;
        }

        const team = res.teams.find((t: any) => t._id === teamId);

        if (!team) {
          this.error = '❌ Aucune équipe correspondante trouvée pour cet ID.';
          this.loading = false;
          return;
        }

        const teamName = team.Name;
        console.log('ℹ️ TeamName à utiliser pour l\'API :', teamName);

        // On récupère les utilisateurs de l’équipe
        this.api.getManagerTeamUsersByName(teamName).subscribe({
          next: (users) => {
            console.log('✅ Utilisateurs récupérés :', users);
            this.users = users;
            this.loading = false;
          },
          error: (err) => {
            console.error('❌ Erreur API manager/team/users :', err);
            this.error = "Impossible de charger les utilisateurs de votre équipe";
            this.loading = false;
          }
        });
      },
      error: (err) => {
        console.error('❌ Erreur API getAllTeams :', err);
        this.error = "Impossible de récupérer les équipes";
        this.loading = false;
      }
    });
  }
}
