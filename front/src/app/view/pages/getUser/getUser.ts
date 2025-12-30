import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
    selector: 'app-users',
    standalone: true,
    imports: [CommonModule],
    template: `
    <h1>Liste des utilisateurs</h1>

    <ul *ngIf="users.length > 0">
      <li *ngFor="let user of users">
        {{ user.firstName }} {{ user.lastName }} - {{ user.role }}
      </li>
    </ul>

    <p *ngIf="users.length === 0">Aucun utilisateur trouvé.</p>
  `
})
export class UsersComponent implements OnInit {

    users: any[] = [];

    constructor(private userService: UserService) { }

    ngOnInit(): void {
        this.userService.getAllUsers().subscribe({
            next: (data) => this.users = data,
            error: (err) => console.error('Erreur chargement users', err)
        });
    }
}
