import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  template: `
  <main>
    <header class="brand-name" style="display: flex; align-items: center; gap: 10px; margin-bottom: 20px;">
      <img class="brand-logo" src="assets/logo.svg" alt="logo" aria-hidden="true" style="height: 100px;">
      <strong style="font-size: 1.5rem;">Bienvenue</strong>
    </header>

    <section class="content" style="align-items: center; text-align: center;">
      <router-outlet></router-outlet>
    </section>
  </main>
  `,
  styleUrls: ['./app.component.css'],
  imports: [RouterOutlet]
})
export class AppComponent {
  title = 'homes';
}