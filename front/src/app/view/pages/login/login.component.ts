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

    this.api.login({ email: this.email, password: this.password }).subscribe({
      next: (res) => {
        const token = (res as any).token;
        if (!token) {
          this.error = 'Réponse invalide du serveur.';
          this.loading = false;
          return;
        }
        localStorage.setItem('access_token', token);
        this.router.navigate(['/home']);
      },
      error: (err) => {
        this.error = err?.error?.message || 'Identifiants invalides.';
        this.loading = false;
      }
    });
  }
}
