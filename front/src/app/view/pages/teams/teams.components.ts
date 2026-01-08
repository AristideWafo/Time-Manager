import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../../services/api.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

interface Team {
  _id: string;
  Name: string;
  NameBeforeEdit?: string;
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
        <input
          type="text"
          [(ngModel)]="newTeamName"
          placeholder="Nom de l'équipe"
        />
        <button
          (click)="createTeam()"
          [disabled]="!newTeamName.trim()"
        >
          Créer
        </button>
      </div>

      <p *ngIf="message" [class.error]="isError" [class.success]="!isError">
        {{ message }}
      </p>

      <hr />

      <!-- Liste des teams -->
      <div class="teams-list">
        <h3>Toutes les équipes</h3>

        <ul>
          <li *ngFor="let team of teams">

            <input type="text" [(ngModel)]="team.Name" />

            <button (click)="updateTeam(team)">
              Modifier
            </button>

            <button
              class="delete-btn"
              (click)="deleteTeam(team)"
            >
              🗑 Supprimer
            </button>

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

  /* ===================== LOAD ===================== */

  loadTeams(): void {
    this.api.getAllTeams().subscribe({
      next: (data: any) => {
        const list = Array.isArray(data) ? data : data.teams || [];
        this.teams = list.map((t: Team) => ({
          ...t,
          NameBeforeEdit: t.Name
        }));
      },
      error: () => {
        this.showMessage('Erreur lors du chargement des équipes ❌', true);
      }
    });
  }

  /* ===================== CREATE ===================== */

  createTeam(): void {
    const name = this.newTeamName.trim();
    if (!name) return;

    this.api.createTeam({ Name: name }).subscribe({
      next: (res: any) => {
        const team = res.team || res;
        this.teams.push({ ...team, NameBeforeEdit: team.Name });
        this.newTeamName = '';
        this.showMessage('Team créée avec succès ✅', false);
      },
      error: (err: any) => {
        if (err.status === 409) {
          this.showMessage('Ce nom de team existe déjà ❌', true);
        } else {
          this.showMessage('Erreur lors de la création ❌', true);
        }
      }
    });
  }

  /* ===================== UPDATE ===================== */

  updateTeam(team: Team): void {
    const newName = team.Name.trim();
    if (!newName) return;

    const payload = {
      CurrentName: team.NameBeforeEdit || team.Name,
      NewName: newName
    };

    this.api.updateTeam(payload).subscribe({
      next: (updated: any) => {
        team.Name = updated.Name || newName;
        team.NameBeforeEdit = team.Name;
        this.showMessage('Nom mis à jour ✅', false);
      },
      error: (err: any) => {
        if (err.status === 409) {
          this.showMessage('Ce nom existe déjà ❌', true);
        } else if (err.status === 404) {
          this.showMessage('Team non trouvée ❌', true);
        } else {
          this.showMessage('Erreur lors de la mise à jour ❌', true);
        }
      }
    });
  }

  /* ===================== DELETE ===================== */

  deleteTeam(team: Team): void {
    const confirmed = confirm(
      `Supprimer définitivement la team "${team.Name}" ?`
    );

    if (!confirmed) return;


    this.api.deleteTeam(team._id).subscribe({
      next: () => {
        this.teams = this.teams.filter(t => t._id !== team._id);
        this.showMessage('Team supprimée ✅', false);
      },
      error: (err: any) => {
        if (err.status === 409) {
          this.showMessage(
            'Impossible de supprimer : des utilisateurs sont encore assignés ❌',
            true
          );
        } else if (err.status === 404) {
          this.showMessage('Team introuvable ❌', true);
        } else {
          this.showMessage('Erreur lors de la suppression ❌', true);
        }
      }
    });
  }

  /* ===================== MESSAGE ===================== */

  private showMessage(msg: string, error: boolean): void {
    this.message = msg;
    this.isError = error;
    setTimeout(() => (this.message = null), 3000);
  }
}
