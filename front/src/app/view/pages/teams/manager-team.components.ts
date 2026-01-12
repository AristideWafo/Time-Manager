import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';
import { Router } from '@angular/router';

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
  <li
    *ngFor="let user of users"
    class="card clickable"
    (click)="goToUser(user._id)"
  >
    <div class="user-info">
      <strong>{{ user.FirstName }} {{ user.LastName }}</strong>
      <span>{{ user.Email }}</span>
      <span
        class="role"
        [ngClass]="user.Role.toLowerCase()"
      >
        {{ user.Role }}
      </span>
    </div>

    <button
      class="view-btn"
      (click)="goToUser(user._id); $event.stopPropagation()"
    >
      👁 Voir les présences
    </button>
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

  constructor(
    private api: ApiService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.loadTeamUsers();
  }

  loadTeamUsers() {
    this.loading = true;
    this.error = null;

    this.api.getManagerTeamUsers().subscribe({
      next: (res: any) => {
        this.users = res.users || [];
        this.loading = false;
      },
      error: (err) => {
        this.error = 'Impossible de charger les utilisateurs de votre équipe';
        this.loading = false;
      }
    });
  }

  goToUser(userId: string) {
    this.router.navigate(['/manager/team/user', userId]);
  }
}
