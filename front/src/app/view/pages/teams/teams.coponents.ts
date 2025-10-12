import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';

@Component({
    selector: 'app-teams',
    standalone: true,
    imports: [CommonModule, FormsModule],
    template: `
  <section class="teams">
    <h1>Gestion des équipes</h1>

    <!-- Formulaire de création -->
    <form (ngSubmit)="createTeam()" #teamForm="ngForm" class="team-form">
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

        <!-- Mode édition -->
        <div *ngIf="team.editing" class="edit-section">
          <input [(ngModel)]="team.name" name="editName" />
          <button (click)="updateTeam(team)">💾 Sauvegarder</button>
          <button (click)="cancelEdit(team)">❌ Annuler</button>
        </div>
      </li>
    </ul>
  </section>
  `,
    styleUrls: ['./teams.component.css']
})
export class TeamsComponent implements OnInit {
    teams: any[] = [];
    newTeamName = '';

    constructor(private api: ApiService) { }

    ngOnInit() {
        this.loadTeams();
    }

    loadTeams() {
        this.api.getAllTeams().subscribe({
            next: (res) => this.teams = res,
            error: (err) => console.error(err)
        });
    }

    createTeam() {
        const data = { name: this.newTeamName };
        this.api.createTeam(data).subscribe({
            next: (res) => {
                this.teams.push(res);
                this.newTeamName = '';
            },
            error: (err) => console.error(err)
        });
    }

    enableEdit(team: any) {
        team.editing = true;
    }

    cancelEdit(team: any) {
        team.editing = false;
        this.loadTeams(); // recharge les données d’origine
    }

    updateTeam(team: any) {
        this.api.updateTeam(team._id, { name: team.name }).subscribe({
            next: () => {
                team.editing = false;
                this.loadTeams();
            },
            error: (err) => console.error(err)
        });
    }

    deleteTeam(id: string) {
        if (!confirm('Supprimer cette équipe ?')) return;
        this.api.deleteTeam(id).subscribe({
            next: () => this.teams = this.teams.filter(t => t._id !== id),
            error: (err) => console.error(err)
        });
    }
}
