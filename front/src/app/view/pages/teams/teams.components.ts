import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';

interface Team {
  _id: string;
  name: string;
  users?: any[];
  editing?: boolean;
}

@Component({
  selector: 'app-teams',
  standalone: true,
  imports: [CommonModule, FormsModule],
  template: `
  <section class="teams">
    <h1>Gestion des équipes</h1>

    <!-- Formulaire de création -->
    <form (ngSubmit)="createTeam()" #teamForm="ngForm">
      <input
        type="text"
        placeholder="Nom de l'équipe"
        name="name"
        [(ngModel)]="newTeamName"
        required
      />
      <button type="submit" [disabled]="!teamForm.valid">Ajouter</button>
    </form>

    <hr>

    <!-- Liste des équipes -->
    <ul>
      <li *ngFor="let team of teams">
        <div *ngIf="!team.editing">
          <strong>{{ team.name }}</strong> ({{ team.users?.length || 0 }} membres)
          <button (click)="enableEdit(team)">✏️ Modifier</button>
          <button (click)="deleteTeam(team._id)">🗑️ Supprimer</button>
        </div>

        <div *ngIf="team.editing">
          <input [(ngModel)]="team.name" name="editName" />
          <button (click)="updateTeam(team)">💾 Sauvegarder</button>
          <button (click)="cancelEdit(team)">❌ Annuler</button>
        </div>
      </li>
    </ul>
  </section>
  `,
  styleUrls: ['./teams.components.css']
})
export class TeamsComponent implements OnInit {
  teams: Team[] = [];
  newTeamName = '';

  constructor(private apiService: ApiService) { }

  ngOnInit() {
    this.loadTeams();
  }

  loadTeams() {
    this.apiService.getAllTeams().subscribe({
      next: (res: any) => this.teams = res,
      error: (err) => console.error('Erreur chargement équipes:', err)
    });
  }

  createTeam() {
    const data = { name: this.newTeamName, users: [] };
    this.apiService.createTeam(data).subscribe({
      next: (res: any) => {
        this.teams.push(res);
        this.newTeamName = '';
      },
      error: (err) => console.error('Erreur création équipe:', err)
    });
  }

  enableEdit(team: Team) {
    team.editing = true;
  }

  cancelEdit(team: Team) {
    team.editing = false;
    this.loadTeams();
  }

  updateTeam(team: Team) {
    const data = { name: team.name, users: team.users || [] };
    this.apiService.updateTeam(team._id, data).subscribe({
      next: () => { team.editing = false; this.loadTeams(); },
      error: (err) => console.error('Erreur mise à jour équipe:', err)
    });
  }

  deleteTeam(id: string) {
    if (!confirm('Supprimer cette équipe ?')) return;
    this.apiService.deleteTeam(id).subscribe({
      next: () => this.teams = this.teams.filter(t => t._id !== id),
      error: (err) => console.error('Erreur suppression équipe:', err)
    });
  }
}
