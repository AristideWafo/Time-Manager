import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  standalone: true,
  template: `
  <section class="home">
    <h1>Accueil</h1>
    <p>Gestion du temps de travail</p>
    <button (click)="goToProfile()">Aller au profil</button>
  </section>
  `,
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(private router: Router) { }
  ngOnInit() {
    console.log("HomeComponent chargé - l'interceptor est actif");
  }
  goToProfile() {
    this.router.navigate(['/profile']);
  }
}