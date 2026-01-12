import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';
import { Router } from '@angular/router';

interface UserProfile {
  firstName: string;
  lastName: string;
  password?: string;
  role: string;
  team: string;
}

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [CommonModule, FormsModule],
  styleUrls: ['./profile.component.css'],
  template: `
    <div class="profile-container">
      <div class="profile-card">
        <div class="profile-header">
          <h1>Mon Profil</h1>

          <div class="profile-actions">
            <button
              class="btn btn-secondary"
              (click)="toggleEditMode()"
              [disabled]="isLoading"
            >
              {{ isEditMode ? 'Annuler' : 'Modifier' }}
            </button>

            <button
              class="btn btn-danger"
              (click)="confirmDelete()"
              [disabled]="isLoading"
            >
              Supprimer le compte
            </button>
          </div>
        </div>

        <div class="profile-content">
          <!-- Vue profil -->
          <div *ngIf="!isEditMode" class="profile-view">
            <div class="profile-info">
              <div>
                <label>Prénom :</label>
                <span>{{ userProfile.firstName }}</span>
              </div>

              <div>
                <label>Nom :</label>
                <span>{{ userProfile.lastName }}</span>
              </div>

              <div>
                <label>Rôle :</label>
                <span>{{ userProfile.role }}</span>
              </div>

              <div>
                <label>Équipe :</label>
                <span>{{ userProfile.team || 'Aucune' }}</span>
              </div>
            </div>
          </div>

          <!-- Formulaire d’édition -->
          <form *ngIf="isEditMode" class="profile-edit" (ngSubmit)="updateProfile()">
            <div>
              <label>Prénom *</label>
              <input [(ngModel)]="editProfile.firstName" name="firstName" required />
            </div>

            <div>
              <label>Nom *</label>
              <input [(ngModel)]="editProfile.lastName" name="lastName" required />
            </div>

            <div>
              <label>Mot de passe *</label>
              <input [(ngModel)]="editProfile.password" name="password" type="password" required />
            </div>

            <div class="profile-actions">
              <button type="submit" class="primary" [disabled]="isLoading">
                {{ isLoading ? 'Sauvegarde...' : 'Sauvegarder' }}
              </button>

              <button type="button" class="secondary" (click)="toggleEditMode()" [disabled]="isLoading">
                Annuler
              </button>
            </div>
          </form>

          <div *ngIf="message" [ngClass]="messageType">
            {{ message }}
          </div>
        </div>

        <!-- Modal de suppression -->
        <div *ngIf="showDeleteModal" class="modal-overlay" (click)="cancelDelete()">
          <div class="modal" (click)="$event.stopPropagation()">
            <h3>Confirmer la suppression</h3>
            <p>Êtes-vous sûr de vouloir supprimer votre compte ?</p>

      

            <button (click)="cancelDelete()" [disabled]="isLoading">
              Annuler
            </button>
          </div>
        </div>
      </div>
    </div>
  `
})
export class ProfileComponent implements OnInit {
  userProfile: UserProfile = {
    firstName: '',
    lastName: '',
    password: '',
    role: '',
    team: ''
  };

  editProfile: UserProfile = { ...this.userProfile };
  isEditMode = false;
  isLoading = false;
  message = '';
  messageType: 'success' | 'error' = 'success';
  showDeleteModal = false;

  constructor(private apiService: ApiService, private router: Router) { }

  ngOnInit() {
    const token = localStorage.getItem('access_token');
    if (!token) {
      this.router.navigate(['/login']);
      return;
    }

    this.loadProfile();
  }

  loadProfile() {
    this.isLoading = true;
    this.apiService.getUser().subscribe({
      next: (res: any) => {
        const u = res.user || {};
        this.userProfile = {
          firstName: u.FirstName || '',
          lastName: u.LastName || '',
          role: u.Role || '',
          team: u.Team || ''
        };
        this.editProfile = { ...this.userProfile, password: '' };
        this.isLoading = false;
      },
      error: () => {
        this.showMessage('Impossible de charger le profil.', 'error');
        this.isLoading = false;
      }
    });
  }

  toggleEditMode() {
    this.isEditMode = !this.isEditMode;
    if (this.isEditMode) this.editProfile = { ...this.userProfile, password: '' };
    this.clearMessage();
  }

  updateProfile() {
    this.isLoading = true;

    const updateData = {
      FirstName: this.editProfile.firstName,
      LastName: this.editProfile.lastName,
      Password: this.editProfile.password,
      Role: this.userProfile.role,
      Team: this.userProfile.team
    };


    this.apiService.updateUser(updateData).subscribe({
      next: () => {
        this.userProfile = { ...this.editProfile };
        this.isEditMode = false;
        this.isLoading = false;
        this.showMessage('Profil mis à jour avec succès !', 'success');
      },
      error: (err) => {
        this.isLoading = false;
        console.error('Erreur update', err);
        this.showMessage('Erreur lors de la mise à jour.', 'error');
      }
    });
  }

  confirmDelete() { this.showDeleteModal = true; }
  cancelDelete() { this.showDeleteModal = false; }

  private showMessage(text: string, type: 'success' | 'error') {
    this.message = text;
    this.messageType = type;
    setTimeout(() => this.clearMessage(), 5000);
  }

  private clearMessage() { this.message = ''; }
}