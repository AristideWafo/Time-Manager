import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-users',
  standalone: true,
  imports: [CommonModule],
  template: `
    <section class="users-page">
      <h1>Liste des utilisateurs</h1>

      <p *ngIf="loading">Chargement...</p>

      <p *ngIf="!loading && users.length === 0">
        ⚠️ Aucun utilisateur trouvé
      </p>

      <table *ngIf="!loading && users.length > 0" class="users-table">
        <thead>
          <tr>
            <th>#</th>
            <th>Nom</th>
            <th>Email</th>
            <th>Rôle</th>
            <th>Équipe</th>
            <th>Modifier équipe</th>
          </tr>
        </thead>

        <tbody>
          <tr *ngFor="let user of users; let i = index">
            <td>{{ i + 1 }}</td>

            <td>{{ user.FirstName }} {{ user.LastName }}</td>

            <td>{{ user.Email }}</td>

            <td>
              <span class="role" [ngClass]="user.Role.toLowerCase()">
                {{ user.Role }}
              </span>
            </td>

            <td>
              {{ user.teamName || '—' }}
            </td>

            <td>
              <select
                [value]="user.teamId"
                (change)="onChangeTeam(user._id, $event)"
              >
                <option value="">— Aucune —</option>

                <option
                  *ngFor="let team of teams"
                  [value]="team._id"
                >
                  {{ team.Name }}
                </option>
              </select>
            </td>
          </tr>
        </tbody>
      </table>
    </section>
  `,
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  users: any[] = [];
  teams: any[] = [];
  loading = true;

  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    console.log('🟢 UsersComponent init');
    this.loadTeamsAndUsers();
  }

  /* ===================== LOAD ===================== */

  private loadTeamsAndUsers(): void {
    console.log('➡️ Appel API getAllTeams()');

    this.apiService.getAllTeams().subscribe({
      next: (res: any) => {
        this.teams = res?.teams ?? [];
        console.log('✅ Teams chargées :', this.teams);
        this.loadUsers();
      },
      error: (err) => {
        console.error('❌ Erreur chargement teams', err);
        this.loading = false;
      }
    });
  }

  private loadUsers(): void {
    console.log('➡️ Appel API getAllUsers()');

    this.apiService.getAllUsers().subscribe({
      next: (data: any) => {
        console.log('✅ Réponse brute API users :', data);

        const rawUsers = data?.users ?? [];

        // 🔥 NORMALISATION ROBUSTE
        this.users = rawUsers.map((u: any) => {
          let teamId = '';
          let teamName = '';

          if (typeof u.Team === 'string') {
            teamId = u.Team;
            const found = this.teams.find(t => t._id === teamId);
            teamName = found?.Name ?? '';
          } else if (u.Team && typeof u.Team === 'object') {
            teamId = u.Team._id;
            teamName = u.Team.Name;
          }

          return {
            ...u,
            teamId,
            teamName
          };
        });

        console.log('📊 users normalisés :', this.users);
        this.loading = false;
      },
      error: (err) => {
        console.error('❌ Erreur API getAllUsers()', err);
        this.loading = false;
      }
    });
  }

  /* ===================== UPDATE TEAM ===================== */

  onChangeTeam(userId: string, event: Event): void {
    const teamId = (event.target as HTMLSelectElement).value;

    console.log(`🔁 Update team user ${userId} → ${teamId}`);

    this.apiService.updateUserTeam(userId, { TeamId: teamId }).subscribe({
      next: () => {
        console.log('✅ Équipe mise à jour');

        const user = this.users.find(u => u._id === userId);
        const team = this.teams.find(t => t._id === teamId);

        if (user) {
          user.teamId = teamId;
          user.teamName = team?.Name ?? '';
        }
      },
      error: (err) => {
        console.error('❌ Erreur update team', err);
      }
    });
  }
}
