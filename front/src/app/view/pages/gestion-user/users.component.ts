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

      <p *ngIf="loading" class="info-message">⏳ Chargement...</p>

      <p *ngIf="!loading && users.length === 0" class="error-message">
        ⚠️ Aucun utilisateur trouvé
      </p>

      <div *ngIf="!loading && users.length > 0" class="users-grid">
        <div *ngFor="let user of users" class="user-card">

          <!-- HEADER -->
          <div class="user-header">
            <strong>{{ user.FirstName }} {{ user.LastName }}</strong>
            <span class="role" [ngClass]="user.Role.toLowerCase()">
              {{ user.Role }}
            </span>
          </div>

          <!-- BODY -->
          <div class="user-body">
            <p>Email : {{ user.Email }}</p>
            <p>Équipe : {{ user.teamName || '—' }}</p>
          </div>

          <!-- FOOTER -->
          <div class="user-footer">
            <select
              [value]="user.teamId"
              (change)="onChangeTeam(user._id, $event)"
            >
              <option value="">— Aucune —</option>
              <option *ngFor="let team of teams" [value]="team._id">
                {{ team.Name }}
              </option>
            </select>

            <button class="delete-btn" (click)="onDeleteUser(user)">
              🗑 Supprimer
            </button>
          </div>

        </div>
      </div>
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
    this.loadTeamsAndUsers();
  }

  /* ===================== LOAD ===================== */

  private loadTeamsAndUsers(): void {
    this.apiService.getAllTeams().subscribe({
      next: (res: any) => {
        this.teams = res?.teams ?? [];
        this.loadUsers();
      },
      error: () => {
        this.loading = false;
      }
    });
  }

  private loadUsers(): void {
    this.apiService.getAllUsers().subscribe({
      next: (res: any) => {
        const rawUsers = res?.users ?? [];

        this.users = rawUsers.map((u: any) => {
          let teamId = '';
          let teamName = '';

          if (typeof u.Team === 'string') {
            teamId = u.Team;
            const team = this.teams.find(t => t._id === teamId);
            teamName = team?.Name ?? '';
          } else if (u.Team && typeof u.Team === 'object') {
            teamId = u.Team._id;
            teamName = u.Team.Name;
          }

          return { ...u, teamId, teamName };
        });

        this.loading = false;
      },
      error: () => {
        this.loading = false;
      }
    });
  }

  /* ===================== UPDATE TEAM ===================== */

  onChangeTeam(userId: string, event: Event): void {
    const teamId = (event.target as HTMLSelectElement).value;

    this.apiService.updateUserTeam(userId, { TeamId: teamId }).subscribe({
      next: () => {
        const user = this.users.find(u => u._id === userId);
        const team = this.teams.find(t => t._id === teamId);

        if (user) {
          user.teamId = teamId;
          user.teamName = team?.Name ?? '';
        }
      }
    });
  }

  /* ===================== DELETE USER ===================== */

  onDeleteUser(user: any): void {
    const confirmed = confirm(
      `Supprimer définitivement ${user.FirstName} ${user.LastName} ?`
    );

    if (!confirmed) return;

    this.apiService.deleteUser(user._id).subscribe({
      next: () => {
        this.users = this.users.filter(u => u._id !== user._id);
      },
      error: (err) => {
        console.error('❌ Erreur suppression utilisateur', err);
        alert('Erreur lors de la suppression');
      }
    });
  }
}
