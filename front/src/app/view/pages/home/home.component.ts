import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';
import jwt_decode from 'jwt-decode';

interface TokenPayload {
  sub: string; // ID utilisateur dans le JWT
  email?: string;
}

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule],
  template: `
  <section class="dashboard">
    <header class="dashboard-header">
      <h1>Tableau de bord</h1>
      <p>Gestion du temps et suivi des présences</p>
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
    </div>
  </section>
  `,
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  message: string | null = null;

  constructor(
    private router: Router,
    private apiService: ApiService
  ) { }

  ngOnInit() {
    console.log("HomeComponent chargé avec succès");

    const token = localStorage.getItem('access_token');
    if (token) {
      try {
        const decoded = jwt_decode<TokenPayload>(token);
        console.log("Token décodé :", decoded);
      } catch (error) {
        console.error("Erreur lors du décodage du token :", error);
      }
    } else {
      console.warn("Aucun token trouvé dans le localStorage.");
    }
  }

  goToProfile() {
    const userId = this.getUserIdFromToken();
    if (!userId) return;
    this.router.navigate(['/profile']);
  }

  goToTeams() {
    this.router.navigate(['/teams']);
  }

  private getUserIdFromToken(): string | null {
    const token = localStorage.getItem('access_token');
    if (!token) return null;

    try {
      const decoded = jwt_decode<TokenPayload>(token);
      return decoded.sub || null;
    } catch {
      return null;
    }
  }

  pointer() {
    const userId = this.getUserIdFromToken();
    if (!userId) {
      this.message = "Erreur : utilisateur non authentifié.";
      setTimeout(() => this.message = null, 3000);
      return;
    }

    const presenceData = {
      presence: {
        Type: "Work",
        Timestamp: new Date().toISOString()
      }
    };

    this.apiService.postPresence(presenceData).subscribe({
      next: () => {
        this.message = "Présence enregistrée avec succès.";
        setTimeout(() => this.message = null, 3000);
      },
      error: (err: any) => {
        console.error("Erreur lors de l'enregistrement :", err);
        this.message = "Erreur lors de l'enregistrement.";
        setTimeout(() => this.message = null, 3000);
      }
    });
  }
}
