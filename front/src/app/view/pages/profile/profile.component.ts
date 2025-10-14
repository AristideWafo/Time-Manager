import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../../services/api.service';
import { Router } from '@angular/router';
import * as jwtDecodeModule from 'jwt-decode';

interface UserProfile {
  id?: string;
  firstName: string;
  lastName: string;
  email: string;
  role: string;
  team: string;
}

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [CommonModule, FormsModule],
  styleUrls: ['./profile.component.css'],
  template: `
    <section class="profile-container">
      <div class="profile-card">
        <header class="profile-header">
          <div>
            <h1>Mon Profil</h1>
            <p>Consultez et gérez vos informations personnelles</p>
          </div>

          <div class="profile-actions">
            <button class="btn btn-secondary" (click)="toggleEditMode()" [disabled]="isLoading">
              {{ isEditMode ? 'Annuler' : 'Modifier' }}
            </button>
            <button class="btn btn-danger" (click)="confirmDelete()" [disabled]="isLoading">
              Supprimer le compte
            </button>
          </div>
        </header>

        <div class="profile-content">
          <!-- Vue profil -->
          <div *ngIf="!isEditMode" class="profile-view">
            <div class="info-row">
              <label>Prénom</label>
              <span>{{ userProfile.firstName }}</span>
            </div>

            <div class="info-row">
              <label>Nom</label>
              <span>{{ userProfile.lastName }}</span>
            </div>

            <div class="info-row">
              <label>Email</label>
              <span>{{ userProfile.email }}</span>
            </div>

            <div class="info-row">
              <label>Rôle</label>
              <span>{{ userProfile.role }}</span>
            </div>

            <div class="info-row">
              <label>Équipe</label>
              <span>{{ userProfile.team || 'Aucune' }}</span>
            </div>
          </div>

          <!-- Formulaire édition -->
          <form *ngIf="isEditMode" (ngSubmit)="updateProfile()" class="profile-form">
            <div class="form-group">
              <label>Prénom *</label>
              <input [(ngModel)]="editProfile.firstName" name="firstName" required />
            </div>

            <div class="form-group">
              <label>Nom *</label>
              <input [(ngModel)]="editProfile.lastName" name="lastName" required />
            </div>

            <div class="form-group">
              <label>Email *</label>
              <input [(ngModel)]="editProfile.email" name="email" required />
            </div>

            <div class="form-actions">
              <button type="submit" class="btn btn-primary" [disabled]="isLoading">
                {{ isLoading ? 'Sauvegarde...' : 'Sauvegarder' }}
              </button>
              <button type="button" class="btn btn-secondary" (click)="toggleEditMode()" [disabled]="isLoading">
                Annuler
              </button>
            </div>
          </form>

          <div *ngIf="message" class="message" [ngClass]="messageType">
            {{ message }}
          </div>
        </div>
      </div>

      <!-- Modal suppression -->
      <div *ngIf="showDeleteModal" class="modal-overlay" (click)="cancelDelete()">
        <div class="modal" (click)="$event.stopPropagation()">
          <h3>Confirmer la suppression</h3>
          <p>Êtes-vous sûr de vouloir supprimer votre compte ?</p>

          <div class="modal-actions">
            <button class="btn btn-danger" (click)="deleteProfile()" [disabled]="isLoading">
              {{ isLoading ? 'Suppression...' : 'Supprimer définitivement' }}
            </button>
            <button class="btn btn-secondary" (click)="cancelDelete()" [disabled]="isLoading">
              Annuler
            </button>
          </div>
        </div>
      </div>
    </section>
  `
})
export class ProfileComponent implements OnInit {
  userProfile: UserProfile = { firstName: '', lastName: '', email: '', role: '', team: '' };
  editProfile: UserProfile = { ...this.userProfile };
  isEditMode = false;
  isLoading = false;
  message = '';
  messageType: 'success' | 'error' = 'success';
  showDeleteModal = false;
  userId!: string;

  constructor(private apiService: ApiService, private router: Router) { }

  ngOnInit() {
    const token = localStorage.getItem('access_token');
    if (!token) {
      this.router.navigate(['/login']);
      return;
    }

    try {
      const decoded: any = (jwtDecodeModule as any).default(token);
      this.userId = decoded.sub;
      this.loadProfile();
    } catch {
      localStorage.removeItem('access_token');
      this.router.navigate(['/login']);
    }
  }

  loadProfile() {
    this.apiService.getUser(this.userId).subscribe({
      next: (res: any) => {
        const user = res.user;
        this.userProfile = {
          firstName: user.FirstName,
          lastName: user.LastName,
          email: user.Email,
          role: user.Role,
          team: user.Team
        };
        this.editProfile = { ...this.userProfile };
      },
      error: () => this.showMessage('Impossible de charger le profil.', 'error')
    });
  }

  toggleEditMode() {
    this.isEditMode = !this.isEditMode;
    if (this.isEditMode) this.editProfile = { ...this.userProfile };
    this.clearMessage();
  }

  updateProfile() {
    if (!this.editProfile.firstName || !this.editProfile.lastName || !this.editProfile.email) {
      this.showMessage('Veuillez remplir tous les champs obligatoires.', 'error');
      return;
    }

    this.isLoading = true;
    this.apiService.updateUser(this.userId, this.editProfile).subscribe({
      next: () => {
        this.userProfile = { ...this.editProfile };
        this.isEditMode = false;
        this.isLoading = false;
        this.showMessage('Profil mis à jour avec succès.', 'success');
      },
      error: () => {
        this.isLoading = false;
        this.showMessage('Erreur lors de la mise à jour.', 'error');
      }
    });
  }

  confirmDelete() {
    this.showDeleteModal = true;
  }

  cancelDelete() {
    this.showDeleteModal = false;
  }

  deleteProfile() {
    this.isLoading = true;
    this.apiService.deleteUser(this.userId).subscribe({
      next: () => {
        this.isLoading = false;
        this.showDeleteModal = false;
        localStorage.removeItem('access_token');
        this.router.navigate(['/login']);
      },
      error: () => {
        this.isLoading = false;
        this.showDeleteModal = false;
        this.showMessage('Impossible de supprimer le profil.', 'error');
      }
    });
  }

  private showMessage(text: string, type: 'success' | 'error') {
    this.message = text;
    this.messageType = type;
    setTimeout(() => this.clearMessage(), 4000);
  }

  private clearMessage() {
    this.message = '';
  }
}
