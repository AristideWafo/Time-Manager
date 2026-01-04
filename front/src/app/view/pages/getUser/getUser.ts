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
          Rôle : {{ user.Role }}<br>
          Équipe : {{ user.Team || '—' }}
        </li>
      </ul>
    </section>
  `,
  styles: [`
    /* --- Page utilisateurs --- */
    .users-page {
        max-width: 900px;
        margin: 0 auto;
        padding: 20px;
        font-family: Arial, sans-serif;
        color: #0a2540;
    }

    .users-page h1 {
        font-size: 1.8rem;
        margin-bottom: 20px;
        text-align: center;
    }

    /* --- Messages --- */
    .users-page p {
        font-size: 1rem;
        text-align: center;
        margin-bottom: 16px;
    }

    .users-page p:first-of-type {
        color: #0d6efd; /* bleu pour le loading */
    }

    .users-page p:last-of-type {
        color: #842029; /* rouge pour aucun utilisateur */
        background-color: #f8d7da;
        border: 1px solid #f5c2c7;
        padding: 10px 14px;
        border-radius: 8px;
    }

    /* --- Liste utilisateurs --- */
    .users-page ul {
        list-style: none;
        padding: 0;
        margin: 0;
        display: flex;
        flex-wrap: wrap;
        gap: 16px;
        justify-content: center;
    }

    .users-page li {
        background-color: #ffffff;
        border: 1px solid rgba(0, 0, 0, 0.05);
        border-radius: 12px;
        box-shadow: 0 4px 12px rgba(0,0,0,0.05);
        padding: 20px;
        flex: 1 1 220px; /* minimum 220px, s’adapte */
        max-width: 250px;
        transition: transform 0.25s ease, box-shadow 0.25s ease;
    }

    .users-page li:hover {
        transform: translateY(-4px);
        box-shadow: 0 8px 20px rgba(0,0,0,0.1);
    }

    .users-page li strong {
        color: #007bff;
        font-size: 1rem;
    }

    .users-page li br {
        margin-bottom: 4px;
    }

    /* --- Responsive --- */
    @media (max-width: 600px) {
        .users-page ul {
            flex-direction: column;
            align-items: center;
        }

        .users-page li {
            max-width: 90%;
        }
    }
  `]
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
