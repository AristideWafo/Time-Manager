import { Routes } from '@angular/router';
import { LoginComponent } from '../pages/login/login.component';
import { HomeComponent } from '../pages/home/home.component';
import { ProfileComponent } from '../pages/profile/profile.component';
import { TeamsComponent } from '../pages/teams/teams.components';
import { PresenceStatsComponent } from '../pages/presence-stats/presence_stats.component';
import { authGuard } from '../guards/auth.guard';

export const routes: Routes = [
    { path: '', redirectTo: 'login', pathMatch: 'full' },
    { path: 'login', component: LoginComponent },
    { path: 'home', component: HomeComponent, canActivate: [authGuard] },
    { path: 'profile', component: ProfileComponent, canActivate: [authGuard] },
    { path: 'teams', component: TeamsComponent, canActivate: [authGuard] },
    { path: 'presence-stats', component: PresenceStatsComponent, canActivate: [authGuard] },
    { path: '**', redirectTo: 'login' }
];


