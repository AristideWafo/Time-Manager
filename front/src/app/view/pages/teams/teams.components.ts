import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../../services/api.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

interface Team {
  _id: string;
  Name: string;
  NameBeforeEdit?: string; // 🔹 utile pour update
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
  styleUrls: ['./teams.components.css']
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

  loadTeams(): void {
    this.api.getAllTeams().subscribe({
      next: (data: any) => {
        const list = Array.isArray(data) ? data : data.teams || [];
        this.teams = list.map((t: Team) => ({ ...t, NameBeforeEdit: t.Name }));
        console.log('Teams API response', this.teams);
      },
      error: (err) => {
        console.error('Erreur chargement teams', err);
        this.showMessage('Erreur lors du chargement des équipes ❌', true);
      }
    });
  }

  createTeam(): void {
    const name = this.newTeamName.trim();
    if (!name) return;

    this.api.createTeam({ Name: name }).subscribe({
      next: (team: any) => {
        const newTeam = team.team || team;
        this.teams.push({ ...newTeam, NameBeforeEdit: newTeam.Name });
        this.newTeamName = '';
        this.showMessage('Team créée avec succès ✅', false);
        console.log('Team créée', newTeam);
      },
      error: (err: any) => {
        console.error('Erreur création team', err);
        if (err.error?.error === 'new name already exists') {
          this.showMessage('Erreur : ce nom de team existe déjà ❌', true);
        } else {
          this.showMessage('Erreur lors de la création de la team ❌', true);
        }
      }
    });
  }

  updateTeam(team: Team): void {
    const newName = team.Name.trim();
    if (!newName) return;

    const payload = { CurrentName: team.NameBeforeEdit || team.Name, NewName: newName };

    this.api.updateTeam(payload).subscribe({
      next: (updated: any) => {
        team.Name = updated.Name || newName;
        team.NameBeforeEdit = team.Name;
        this.showMessage('Nom de la team mis à jour ✅', false);
        console.log('Team mise à jour', updated);
      },
      error: (err: any) => {
        console.error('Erreur update team', err);
        if (err.status === 409) {
          this.showMessage('Erreur : ce nom de team existe déjà ❌', true);
        } else if (err.status === 404) {
          this.showMessage('Erreur : équipe non trouvée ❌', true);
        } else {
          this.showMessage('Erreur lors de la mise à jour ❌', true);
        }
      }
    });
  }

  private showMessage(msg: string, error: boolean) {
    this.message = msg;
    this.isError = error;
    setTimeout(() => this.message = null, 3000);
  }
}
