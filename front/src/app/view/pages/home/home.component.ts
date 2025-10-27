import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule],
  template: `
  <section class="dashboard">
    <header class="dashboard-header">
      <h1>Tableau de bord</h1>
      <p>Gestion du temps et suivi des présences</p>

      <!--  Message personnalisé -->
      <div class="welcome-box" *ngIf="firstName && lastName && role">
         Bonjour <strong>{{ firstName }} {{ lastName }}</strong> !
        <br>
        <span class="role">Votre rôle : <strong>{{ role }}</strong></span>
      </div>
    </header>

    <div class="dashboard-cards">
      <!-- Profil -->
      <div class="card">
        <h2>Profil</h2>
        <p>Accédez à vos informations personnelles et à vos paramètres.</p>
        <button (click)="goToProfile()">Voir le profil</button>
      </div>

      <!-- Équipes -->
      <div class="card">
        <h2>Équipes</h2>
        <p>Gérez vos équipes et visualisez les membres.</p>
        <button (click)="goToTeams()">Gérer les équipes</button>
      </div>

      <!-- Présence -->
      <div class="card">
        <h2>Présence</h2>
        <p>Enregistrez votre présence quotidienne.</p>
        <button (click)="pointer()">Pointer</button>
        <p *ngIf="message" class="message">{{ message }}</p>
      </div>

      <!-- Statistiques -->
      <div class="card">
        <h2>Statistiques</h2>
        <p>Consultez vos présences et temps de travail.</p>
        <button (click)="goToStats()">Voir les statistiques</button>
      </div>
    </div>
  </section>
  `,
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  message: string | null = null;
  firstName: string | null = null;
  lastName: string | null = null;
  role: string | null = null;

  constructor(private router: Router, private apiService: ApiService) { }

  ngOnInit() {
    const token = localStorage.getItem('access_token');
    const userId = localStorage.getItem('user_id');

    // ✅ Récupération des infos utilisateur depuis le stockage local
    this.firstName = localStorage.getItem('FirstName');
    this.lastName = localStorage.getItem('LastName');
    this.role = localStorage.getItem('Role');

    if (!token || !userId) {
      console.warn('Token ou ID utilisateur manquant.');
      this.router.navigate(['/login']);
    }
  }

  goToProfile() { this.router.navigate(['/profile']); }
  goToTeams() { this.router.navigate(['/teams']); }
  goToStats() { this.router.navigate(['/presence-stats']); }

  pointer() {
    const userId = localStorage.getItem('user_id');
    if (!userId) {
      this.message = "❌ Impossible de pointer : utilisateur non identifié.";
      return;
    }

    const presenceData = {
      Type: "Work",
      Timestamp: new Date().toISOString(),
      User: userId
    };

    console.log("Envoi des données de présence :", presenceData);

    this.apiService.postPresence(presenceData).subscribe({
      next: () => {
        this.message = "✅ Présence enregistrée avec succès.";
        setTimeout(() => this.message = null, 3000);
      },
      error: (err: any) => {
        console.error("❌ Erreur lors de l'enregistrement :", err);
        this.message = "Erreur lors de l'enregistrement.";
        setTimeout(() => this.message = null, 3000);
      }
    });
  }
}
