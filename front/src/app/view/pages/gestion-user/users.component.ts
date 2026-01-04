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

      <ul *ngIf="!loading && users.length > 0">
        <li *ngFor="let user of users; let i = index">
          <strong>#{{ i + 1 }}</strong><br>
          Prénom : {{ user.FirstName }}<br>
          Nom : {{ user.LastName }}<br>
          Email : {{ user.Email }}<br>
          Rôle : <span class="role" [ngClass]="user.Role.toLowerCase()">{{ user.Role }}</span><br>
          Équipe : {{ user.Team || '—' }}
        </li>
      </ul>
    </section>
  `,
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  users: any[] = [];
  loading = true;

  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    console.log('🟢 UsersComponent init');
    this.loadUsers();
  }

  private loadUsers(): void {
    console.log('➡️ Appel API getAllUsers()');

    this.apiService.getAllUsers().subscribe({
      next: (data: any) => {
        console.log('✅ Réponse brute API :', data);

        this.users = data?.users ?? [];

        console.log('📊 users après extraction :', this.users);
        console.log('📊 users.length =', this.users.length);

        this.loading = false;
      },
      error: (err) => {
        console.error('❌ Erreur API getAllUsers()', err);
        this.loading = false;
      }
    });
  }
}
