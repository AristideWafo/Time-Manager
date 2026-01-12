import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule],
  styleUrls: ['./login.component.css'],
  template: `
  <div class="login-container">
    <div class="login-card">
      <h1>Connexion</h1>
      <form (ngSubmit)="onLogin()" #f="ngForm">
        <label>Email</label>
        <input name="email" type="email" [(ngModel)]="email" placeholder="Votre email" required>

        <label>Mot de passe</label>
        <input name="password" type="password" [(ngModel)]="password" placeholder="Votre mot de passe" required>

        <button class="primary" type="submit" [disabled]="loading || !f.form.valid">
          {{ loading ? 'Connexion...' : 'Se connecter' }}
        </button>

        <div class="error" *ngIf="error">{{ error }}</div>
      </form>
    </div>
  </div>
  `
})
export class LoginComponent {
  email = '';
  password = '';
  loading = false;
  error = '';

  constructor(private router: Router, private api: ApiService) { }

  onLogin() {
    if (this.loading) return;
    this.error = '';
    this.loading = true;

    this.api.login({ Email: this.email, Password: this.password }).subscribe({
      next: (res: any) => {
        if (!res.Token || !res._id) {
          this.error = 'Réponse du serveur invalide.';
          this.loading = false;
          return;
        }

        localStorage.setItem('access_token', res.Token);
        localStorage.setItem('user_id', res._id);
        localStorage.setItem('Role', res.Role);
        localStorage.setItem('FirstName', res.FirstName);
        localStorage.setItem('LastName', res.LastName);
        localStorage.setItem('Team', res.Team);

        if (res.Role === 'ADMIN') {
          this.router.navigate(['/admin-home']);
        }
        else if (res.Role === 'MANAGER') {
          this.router.navigate(['/manager-home']);
        }
        else {
          this.router.navigate(['/employee-home']);
        }

      },
      error: (err) => {
        console.error('❌ Erreur lors de la connexion :', err);
        this.error = err?.error?.message || 'Identifiants invalides.';
        this.loading = false;
      }
    });
  }
}
