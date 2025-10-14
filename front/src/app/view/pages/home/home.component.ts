import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';
import jwt_decode from 'jwt-decode';

interface TokenPayload {
  _id: string;
  email?: string;
}

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule],
  template: `
  <section class="home">
    <h1>Accueil</h1>
    <p>Gestion du temps de travail</p>

    <button (click)="goToProfile()">Aller au profil</button>
    <button (click)="goToTeams()">Gestion des équipes</button>
    <button (click)="pointer()">Pointer</button>

    <p *ngIf="message" class="message">{{ message }}</p>
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
    console.log("HomeComponent chargé - l'interceptor est actif");
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
      const decoded = jwt_decode(token) as TokenPayload;
      return decoded._id || null;
    } catch (e) {
      console.error('Token invalide :', e);
      return null;
    }
  }

  pointer() {
    const userId = this.getUserIdFromToken();
    if (!userId) {
      this.message = "Erreur : utilisateur non authentifié ❌";
      setTimeout(() => this.message = null, 3000);
      return;
    }

    const presenceData = {
      presence: {
        Type: "Default",
        Timestamp: new Date().toISOString()
      }
    };

    this.apiService.postPresence(presenceData).subscribe({
      next: () => {
        this.message = "Présence enregistrée avec succès ✅";
        setTimeout(() => this.message = null, 3000);
      },
      error: (err: any) => {
        console.error('Erreur lors de l\'enregistrement de la présence :', err);
        this.message = "Erreur lors de l'enregistrement ❌";
        setTimeout(() => this.message = null, 3000);
      }
    });
  }
}
