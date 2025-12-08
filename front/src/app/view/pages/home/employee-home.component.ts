import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

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
          <h2>Profil</h2>
          <p>Accédez à vos informations personnelles et à vos paramètres.</p>
          <button (click)="goToProfile()">Voir le profil</button>
        </div>

        <div class="card">
          <h2>Présence</h2>
          <p>Enregistrez votre présence quotidienne.</p>
          <button (click)="pointer()">Pointer</button>
          <p *ngIf="message" class="message">{{ message }}</p>
        </div>

        <div class="card">
          <h2>Statistiques</h2>
          <p>Consultez vos présences et temps de travail.</p>
          <button (click)="goToStats()">Voir les statistiques</button>
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

    constructor(private router: Router, private api: ApiService) { }

    ngOnInit() {
        this.firstName = localStorage.getItem('FirstName');
        this.lastName = localStorage.getItem('LastName');
        this.role = localStorage.getItem('Role');
    }

    goToProfile() { this.router.navigate(['/profile']); }
    goToStats() { this.router.navigate(['/presence-stats']); }

    pointer() {
        const userId = localStorage.getItem('user_id');
        if (!userId) return;

        const data = { Type: 'Work', Timestamp: new Date().toISOString(), User: userId };

        this.api.postPresence(data).subscribe({
            next: () => { this.message = "✅ Présence enregistrée avec succès."; setTimeout(() => this.message = null, 3000); },
            error: () => { this.message = "❌ Erreur lors de l'enregistrement."; setTimeout(() => this.message = null, 3000); }
        });
    }

    logout() {
        localStorage.clear();
        this.router.navigate(['/login']);
    }
}
