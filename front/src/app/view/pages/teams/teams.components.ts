// teams.component.ts
import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../../services/api.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

interface Team {
  _id: string;
  Name: string;
}

@Component({
  selector: 'app-teams',
  standalone: true,
  imports: [CommonModule, FormsModule],
  template: `
    <section class="teams-page">
      <h1>Gestion des Teams</h1>

      <!-- Création de team -->
      <div class="create-team">
        <h3>Créer une nouvelle équipe</h3>
        <input type="text" [(ngModel)]="newTeamName" placeholder="Nom de l'équipe" />
        <button (click)="createTeam()" [disabled]="!newTeamName.trim()">Créer</button>
      </div>

      <p *ngIf="message" [class.error]="isError" [class.success]="!isError">
        {{ message }}
      </p>

      <hr/>

      <!-- Liste des teams -->
      <div class="teams-list">
        <h3>Toutes les équipes</h3>
        <ul>
          <li *ngFor="let team of teams">
            <input type="text" [(ngModel)]="team.Name" />
            <button (click)="updateTeam(team)">Modifier le nom</button>
            <span class="team-id">ID: {{ team._id }}</span>
          </li>
        </ul>
      </div>
    </section>
  `,
  styles: [`
    .teams-page { max-width: 600px; margin: 0 auto; font-family: Arial; }
    input { margin-right: 10px; }
    button { margin-right: 20px; }
    hr { margin: 20px 0; }
    .team-id { color: gray; font-size: 0.9em; }
    .success { color: green; }
    .error { color: red; }
  `]
})
export class TeamsComponent implements OnInit {
  teams: Team[] = [];
  newTeamName = '';
  message: string | null = null;
  isError = false;

  constructor(private api: ApiService) { }

  ngOnInit(): void {
    this.loadTeams();
  }

  // 🔹 Récupération de toutes les teams
  loadTeams(): void {
    this.api.getAllTeams().subscribe({
      next: (data: any) => {
        console.log('Teams API response', data);
        // Forcer un tableau quelle que soit la structure
        this.teams = Array.isArray(data) ? data : data.teams || [];
      },
      error: (err) => {
        console.error('Erreur chargement teams', err);
        this.showMessage('Erreur lors du chargement des équipes ❌', true);
      }
    });
  }

  // 🔹 Création d'une nouvelle team
  createTeam(): void {
    const name = this.newTeamName.trim();
    if (!name) return;

    this.api.createTeam({ Name: name }).subscribe({
      next: (team: any) => {
        // Backend peut renvoyer { team: {...} } ou juste l'objet
        const newTeam = team.team || team;
        if (!Array.isArray(this.teams)) this.teams = [];
        this.teams.push(newTeam);
        this.newTeamName = '';
        this.showMessage('Team créée avec succès ✅', false);
        console.log('Team créée', newTeam);
      },
      error: (err) => {
        console.error('Erreur création team', err);
        this.showMessage('Erreur lors de la création de la team ❌', true);
      }
    });
  }

  // 🔹 Mise à jour du nom d'une team
  updateTeam(team: Team): void {
    const newName = team.Name.trim();
    if (!newName) return;

    this.api.updateTeam({ CurrentName: team.Name, NewName: newName }).subscribe({
      next: (updated: any) => {
        team.Name = updated.Name || newName;
        this.showMessage('Nom de la team mis à jour ✅', false);
        console.log('Team mise à jour', updated);
      },
      error: (err) => {
        console.error('Erreur update team', err);
        this.showMessage('Erreur lors de la mise à jour ❌', true);
      }
    });
  }

  // 🔹 Affichage des messages UI
  private showMessage(msg: string, error: boolean) {
    this.message = msg;
    this.isError = error;
    setTimeout(() => this.message = null, 3000);
  }
}
