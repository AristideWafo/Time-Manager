import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-manager-home',
  standalone: true,
  imports: [CommonModule],
  template: `
    <section class="dashboard">
      <header class="dashboard-header">
        <h1>Tableau de bord</h1>
        <p>Gestion des équipes</p>

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
            <div class="icon">👥</div>
            <h2>Équipes</h2>
          </div>
          <div class="card-body">
            <p>Gérez vos équipes et visualisez les membres.</p>
          </div>
          <div class="card-footer">
            <button (click)="goToTeams()">Voir les équipes</button>
          </div>
        </div>
      </div>
    </section>
  `,
  styleUrls: ['./manager-home.component.css']
})
export class ManagerHomeComponent implements OnInit {
  firstName: string | null = null;
  lastName: string | null = null;
  role: string | null = 'MANAGER';

  constructor(private router: Router) { }

  ngOnInit() {
    this.firstName = localStorage.getItem('FirstName');
    this.lastName = localStorage.getItem('LastName');
    this.role = localStorage.getItem('Role');
  }

  goToProfile() { this.router.navigate(['/profile']); }
  goToTeams() { this.router.navigate(['/manager/team']); }
  logout() { localStorage.clear(); this.router.navigate(['/login']); }
}
