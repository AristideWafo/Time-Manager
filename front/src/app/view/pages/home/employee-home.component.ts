import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

interface Presence {
  Type: string;
  Timestamp: string;
  User: string;
}

@Component({
  selector: 'app-employee-home',
  standalone: true,
  imports: [CommonModule],
  template: `
    <section class="dashboard">
      <header class="dashboard-header">
        <h1>Tableau de bord</h1>
        <p>Gestion du temps et suivi des présences</p>

        <button class="logout-button" (click)="logout()">Déconnexion</button>

        <div class="welcome-box" *ngIf="firstName && lastName && role">
          Bonjour <strong>{{ firstName }} {{ lastName }}</strong> !<br>
          <span class="role">Votre rôle : <strong>{{ role }}</strong></span>
        </div>
      </header>

      <div class="dashboard-cards">
        <div class="card">
          <div class="card-header">
            <div class="icon">👤</div>
            <h2>Profil</h2>
          </div>
          <div class="card-body">
            <p>Accédez à vos informations personnelles et à vos paramètres.</p>
          </div>
          <div class="card-footer">
            <button (click)="goToProfile()">Voir le profil</button>
          </div>
        </div>

        <div class="card">
          <div class="card-header">
            <div class="icon">⏰</div>
            <h2>Présence</h2>
          </div>
          <div class="card-body">
            <p>Enregistrez votre présence quotidienne.</p>
            <p *ngIf="message" class="message">{{ message }}</p>
          </div>
          <div class="card-footer">
            <button (click)="pointer()" [disabled]="isPosting">
              Pointer {{ nextBadgeType }}
            </button>
          </div>
        </div>

        <div class="card">
          <div class="card-header">
            <div class="icon">📊</div>
            <h2>Statistiques</h2>
          </div>
          <div class="card-body">
            <p>Consultez vos présences et temps de travail.</p>
          </div>
          <div class="card-footer">
            <button (click)="goToStats()">Voir les statistiques</button>
          </div>
        </div>
      </div>
    </section>
  `,
  styleUrls: ['./employee-home.component.css']
})
export class EmployeeHomeComponent implements OnInit {
  firstName: string | null = null;
  lastName: string | null = null;
  role: string | null = 'EMPLOYEE';
  message: string | null = null;

  // Liste des présences du jour pour cet utilisateur
  todaysPresences: Presence[] = [];

  // Type du prochain badge
  nextBadgeType: 'Check-in' | 'Check-out' = 'Check-in';

  // Bloquer le bouton pendant l'envoi
  isPosting = false;

  constructor(private router: Router, private api: ApiService) { }

  ngOnInit() {
    this.firstName = localStorage.getItem('FirstName');
    this.lastName = localStorage.getItem('LastName');
    this.role = localStorage.getItem('Role');

    const userId = localStorage.getItem('user_id');
    if (userId) {
      // Récupérer toutes les présences pour déterminer le type du prochain badge
      this.api.getPresence().subscribe({
        next: (res: any) => {
          const allPresences: Presence[] = res.user || [];
          const today = new Date().toISOString().slice(0, 10);

          this.todaysPresences = allPresences.filter(
            p => p.User === userId && p.Timestamp.slice(0, 10) === today
          );

          // Déterminer le type du prochain badge
          this.nextBadgeType = (this.todaysPresences.length % 2 === 0) ? 'Check-in' : 'Check-out';
        },
        error: err => console.error('Erreur récupération présences', err)
      });
    }
  }

  pointer() {
    if (this.isPosting) return; // empêche les clics multiples
    this.isPosting = true;

    const userId = localStorage.getItem('user_id');
    if (!userId) {
      this.isPosting = false;
      return;
    }

    const data: Presence = {
      Type: this.nextBadgeType,
      Timestamp: new Date().toISOString(),
      User: userId
    };

    this.api.postPresence(data).subscribe({
      next: () => {
        this.message = `✅ ${this.nextBadgeType} enregistré avec succès.`;
        // On met à jour la liste pour le prochain badge
        this.todaysPresences.push(data);
        this.nextBadgeType = (this.todaysPresences.length % 2 === 0) ? 'Check-in' : 'Check-out';
        setTimeout(() => this.message = null, 3000);
        this.isPosting = false;
      },
      error: () => {
        this.message = `❌ Erreur lors de l'enregistrement.`;
        setTimeout(() => this.message = null, 3000);
        this.isPosting = false;
      }
    });
  }

  goToProfile() { this.router.navigate(['/profile']); }
  goToStats() { this.router.navigate(['/presence-stats']); }

  logout() {
    localStorage.clear();
    this.router.navigate(['/login']);
  }
}
