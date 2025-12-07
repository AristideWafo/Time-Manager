import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';

interface Team {
  _id: string;
  Name: string;
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
        name="Name"
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
          <strong>{{ team.Name }}</strong>
          <button (click)="enableEdit(team)">✏️ Modifier</button>
          <button (click)="deleteTeam(team._id)">🗑️ Supprimer</button>

          <!-- 🔥 Afficher les membres de la team -->
          <div class="members">
            <h4>Membres :</h4>
            <ul>
              <li *ngFor="let m of getMembers(team._id)">
                {{ m.FirstName }} {{ m.LastName }} ({{ m.Role }})
              </li>
            </ul>
          </div>
        </div>

        <div *ngIf="team.editing">
          <input [(ngModel)]="team.Name" name="editName" required />
          <button (click)="updateTeam(team)"> Sauvegarder</button>
          <button (click)="cancelEdit(team)"> Annuler</button>
        </div>
      </li>
    </ul>
  </section>
  `,
  styleUrls: ['./teams.components.css']
})
export class TeamsComponent implements OnInit {

  teams: Team[] = [];
  users: any[] = []; // 👉 Tous les users pour filtrer ensuite
  newTeamName = '';

  constructor(private apiService: ApiService) { }

  ngOnInit() {
    this.loadTeams();
    this.loadUsers(); // 👉 charge tous les users une fois
  }

  loadTeams() {
    this.apiService.getAllTeams().subscribe({
      next: (res: any[]) => this.teams = res,
      error: (err) => console.error('Erreur chargement équipes:', err)
    });
  }

  loadUsers() {
    this.apiService.getAllUsers().subscribe({
      next: (res: any[]) => this.users = res,
      error: (err) => console.error('Erreur chargement users:', err)
    });
  }

  /** Renvoie les utilisateurs appartenant à cette équipe */
  getMembers(teamId: string) {
    return this.users.filter(u => u.Team === teamId);
  }

  createTeam() {
    const data = { Name: this.newTeamName };

    this.apiService.createTeam(data).subscribe({
      next: (team: any) => {
        this.teams.push(team);
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
    const data = { Name: team.Name };

    this.apiService.updateTeam(team._id, data).subscribe({
      next: () => {
        team.editing = false;
        this.loadTeams();
      },
      error: (err) => console.error('Erreur mise à jour équipe:', err)
    });
  }

  deleteTeam(id: string) {
    if (!confirm('Supprimer cette équipe ?')) return;

    this.apiService.deleteTeam(id).subscribe({
      next: () => {
        this.teams = this.teams.filter(t => t._id !== id);
        this.users = this.users.filter(u => u.Team !== id); // supprime les membres du cache
      },
      error: (err) => console.error('Erreur suppression équipe:', err)
    });
  }

}
