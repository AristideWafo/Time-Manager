import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../../services/api.service';

@Component({
  selector: 'app-users',
  standalone: true,
  imports: [CommonModule],
  template: `
      <h1>Liste des utilisateurs</h1>

      <ul *ngIf="users.length > 0; else noUsers">
        <li *ngFor="let user of users">
          {{ user.firstName }} {{ user.lastName }} - {{ user.role }}
        </li>
      </ul>

      <ng-template #noUsers>
        <p>Aucun utilisateur trouvé.</p>
      </ng-template>
    `
})
export class UsersComponent implements OnInit {

  users: any[] = [];

  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    this.apiService.getAllUsers().subscribe({
      next: (data) => this.users = data,
      error: (err) => console.error('Erreur chargement users', err)
    });
  }
}
