import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  template: `
  <div class="login-container">
    <div class="login-card">
      <h1>Connexion</h1>
      <form>
        <label>Email</label>
        <input type="email" placeholder="email@exemple.com">
      
        <label>Mot de passe</label>
        <input type="password" placeholder="Votre mot de passe">

        <button class="primary" type="button" (click)="onLogin()">Se connecter</button>
      </form>
    </div>
  </div>
  `,
  styles: [`
  .login-container {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #f6f7fb;
  }
  .login-card {
    width: 360px;
    background: white;
    border-radius: 12px;
    padding: 24px;
    box-shadow: 0 10px 30px rgba(0,0,0,0.08);
  }
  h1 {
    margin: 0 0 16px;
    font-size: 24px;
  }
  form {
    display: grid;
    gap: 12px;
  }
  label {
    font-size: 12px;
    color: #555;
  }
  input {
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid #ddd;
    outline: none;
  }
  input:focus {
    border-color: var(--primary-color);
    box-shadow: 0 0 0 3px color-mix(in srgb, var(--primary-color) 20%, transparent);
  }
  .primary {
    padding: 10px 14px;
    border-radius: 8px;
    border: 1px solid var(--primary-color);
    background: var(--primary-color);
    color: #fff;
    cursor: pointer;
  }
  `]
})
export class LoginComponent {
  constructor(private router: Router) { }

  onLogin() {
    localStorage.setItem('access_token', 'FAKE_TOKEN_DEV_123');
    this.router.navigate(['/home']);
  }
}


