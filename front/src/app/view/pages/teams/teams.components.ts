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
    <h1>Votre équipe</h1>

    <hr>

    <div *ngIf="teamName; else noTeam">
      <h2>{{ teamName }}</h2>

      <h3>Membres :</h3>
      <ul>
        <li *ngFor="let user of members">
          {{ user.FirstName }} {{ user.LastName }} ({{ user.Role }})
        </li>
      </ul>

      <p *ngIf="members.length === 0">Aucun membre trouvé.</p>
    </div>

    <ng-template #noTeam>
      <p>Aucune équipe associée à votre compte.</p>
    </ng-template>

  </section>
  `,
  styleUrls: ['./teams.components.css']
})
export class TeamsComponent implements OnInit {

  teamName: string | null = null;
  members: any[] = [];

  constructor(private apiService: ApiService) { }

  ngOnInit() {
    this.loadUserTeam();
  }

  loadUserTeam() {
    this.apiService.getUser().subscribe({
      next: (user: any) => {
        console.log("User connecté :", user);


        this.teamName = user.Team || null;

        if (this.teamName) {
          this.loadMembers(this.teamName);
        }
      },
      error: (err) => {
        console.error('Erreur récupération user:', err);
      }
    });
  }

  loadMembers(teamName: string) {
    this.apiService.getTeamUsers(teamName).subscribe({
      next: (res: any) => {
        console.log("Réponse API membres :", res);


        this.members = res.team || [];
      },
      error: (err) => {
        console.error('Erreur chargement membres:', err);
      }
    });
  }

}
