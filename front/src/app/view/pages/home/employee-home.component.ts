import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

interface Presence {
  Type: 'Check-in' | 'Check-out';
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
            <p>Accédez à vos informations personnelles.</p>
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
            <p>Enregistrez votre présence.</p>
            <p *ngIf="message" class="message">{{ message }}</p>
          </div>
          <div class="card-footer">
            <button (click)="pointer()" [disabled]="isPosting">
              Pointer {{ allowedNextBadge }}
            </button>
          </div>
        </div>

        <div class="card">
          <div class="card-header">
            <div class="icon">📊</div>
            <h2>Statistiques</h2>
          </div>
          <div class="card-body">
            <p>Consultez vos temps de travail.</p>
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
  isPosting = false;

  /** Liste complète des présences reçues (tous les jours, sans filtrage) */
  allPresences: Presence[] = [];

  /** Badge autorisé pour le prochain clic */
  allowedNextBadge: 'Check-in' | 'Check-out' = 'Check-in';

  constructor(
    private router: Router,
    private api: ApiService
  ) { }

  ngOnInit(): void {
    this.firstName = localStorage.getItem('FirstName');
    this.lastName = localStorage.getItem('LastName');
    this.role = localStorage.getItem('Role');

    this.loadPresences();
  }

  // =========================
  // 🔁 CHARGEMENT DES PRÉSENCES
  // =========================
  private loadPresences(): void {
    const userId = localStorage.getItem('user_id');
    if (!userId) return;

    console.log(`🔁 loadPresences | userId = ${userId}`);

    this.api.getPresence().subscribe({
      next: (res: any) => {
        // ✅ On prend toutes les présences telles qu'elles sont
        this.allPresences = res.user || [];
        console.log('📦 Toutes les présences reçues: ', this.allPresences);

        // Calcul du prochain badge à partir du dernier
        this.allowedNextBadge = this.computeAllowedNextBadge();
        console.log('🧠 Résultat final attendu: ', this.allowedNextBadge);
      },
      error: err => console.error('❌ Erreur récupération présences', err)
    });
  }

  // =========================
  // 🧠 CALCUL DU PROCHAIN BADGE (dernier pointage)
  // =========================
  private computeAllowedNextBadge(): 'Check-in' | 'Check-out' {
    if (this.allPresences.length === 0) return 'Check-in';

    const last = this.allPresences[this.allPresences.length - 1];
    const next = last.Type === 'Check-in' ? 'Check-out' : 'Check-in';
    console.log('➡️ Action autorisée (dernier badge): ', next);
    return next;
  }

  // =========================
  // ⛔ POST PROTÉGÉ
  // =========================
  pointer(): void {
    if (this.isPosting) return;
    this.isPosting = true;

    const userId = localStorage.getItem('user_id');
    if (!userId) {
      this.isPosting = false;
      return;
    }

    const badgeType = this.allowedNextBadge;

    const data: Presence = {
      Type: badgeType,
      Timestamp: new Date().toISOString(),
      User: userId
    };

    console.log('👉 CLICK pointer()');
    console.log('📤 Badge envoyé: ', badgeType);

    this.api.postPresence(data).subscribe({
      next: () => {
        this.message = `✅ ${badgeType} enregistré.`;

        // 🔁 Relecture complète → cohérence garantie
        this.loadPresences();

        setTimeout(() => this.message = null, 3000);
        this.isPosting = false;
        console.log('✅ POST OK');
      },
      error: () => {
        this.message = '❌ Erreur lors de l’enregistrement.';
        setTimeout(() => this.message = null, 3000);
        this.isPosting = false;
      }
    });
  }

  // =========================
  // NAVIGATION
  // =========================
  goToProfile(): void {
    this.router.navigate(['/profile']);
  }

  goToStats(): void {
    this.router.navigate(['/presence-stats']);
  }

  logout(): void {
    localStorage.clear();
    this.router.navigate(['/login']);
  }
}
