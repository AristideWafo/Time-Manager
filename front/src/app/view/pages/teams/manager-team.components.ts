import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-manager-team',
  standalone: true,
  imports: [CommonModule],
  template: `
    <section class="manager-team-page">
      <h1>Utilisateurs de votre équipe</h1>

      <div *ngIf="loading" class="message success">Chargement...</div>
      <div *ngIf="error" class="message error">{{ error }}</div>

      <ul *ngIf="!loading && users.length > 0">
        <li *ngFor="let user of users" class="card">
          <strong>{{ user.FirstName }} {{ user.LastName }}</strong>
          <span>{{ user.Email }}</span>
          <span class="role" [ngClass]="user.Role.toLowerCase()">{{ user.Role }}</span>
        </li>
      </ul>

      <div *ngIf="users.length === 0 && !loading" class="message error">
        Aucun utilisateur trouvé.
      </div>
    </section>
  `,
  styleUrls: ['./manager-team.component.css']
})
export class ManagerTeamComponent implements OnInit {

  users: any[] = [];
  loading = false;
  error: string | null = null;

  constructor(private api: ApiService) { }

  ngOnInit(): void {
    console.log('🟢 ManagerTeamComponent init');
    this.loadTeamUsers();
  }

  loadTeamUsers() {
    this.loading = true;
    this.error = null;

    this.api.getManagerTeamUsers().subscribe({
      next: (res: any) => {
        console.log('✅ Utilisateurs récupérés :', res);
        this.users = res.users || [];
        this.loading = false;
      },
      error: (err) => {
        console.error('❌ Erreur API manager/team/users :', err);
        this.error = 'Impossible de charger les utilisateurs de votre équipe';
        this.loading = false;
      }
    });
  }
}
