import { Routes } from '@angular/router';
import { LoginComponent } from '../pages/login/login.component';
import { ProfileComponent } from '../pages/profile/profile.component';
import { TeamsComponent } from '../pages/teams/teams.components';
import { PresenceStatsComponent } from '../pages/presence-stats/presence_stats.component';
import { CreateUserComponent } from '../pages/create-user/create_user.component';

import { authGuard } from '../guards/auth.guard';
import { roleGuard } from '../guards/role.guard';

import { EmployeeHomeComponent } from '../pages/home/employee-home.component';
import { ManagerHomeComponent } from '../pages/home/manager-home.component';
import { AdminHomeComponent } from '../pages/home/admin-home.component';

import { UsersComponent } from '../pages/getUser/getUser';

export const routes: Routes = [
    { path: '', redirectTo: 'login', pathMatch: 'full' },

    { path: 'login', component: LoginComponent },

    // EMPLOYEE
    {
        path: 'employee-home',
        component: EmployeeHomeComponent,
        canActivate: [authGuard, roleGuard('EMPLOYEE')]
    },

    // MANAGER
    {
        path: 'manager-home',
        component: ManagerHomeComponent,
        canActivate: [authGuard, roleGuard('MANAGER')]
    },

    {
        path: 'manager/team',
        loadComponent: () =>
            import('../pages/teams/manager-team.components')
                .then(m => m.ManagerTeamComponent),
        canActivate: [authGuard, roleGuard('MANAGER')]
    },


    // ADMIN
    {
        path: 'admin-home',
        component: AdminHomeComponent,
        canActivate: [authGuard, roleGuard('ADMIN')]
    },

    // Commun
    { path: 'profile', component: ProfileComponent, canActivate: [authGuard] },
    { path: 'teams', component: TeamsComponent, canActivate: [authGuard] },
    { path: 'presence-stats', component: PresenceStatsComponent, canActivate: [authGuard] },

    // ADMIN only
    { path: 'create-user', component: CreateUserComponent, canActivate: [authGuard, roleGuard('ADMIN')] },
    {
        path: 'users',
        component: UsersComponent,
        canActivate: [authGuard, roleGuard('ADMIN')]
    },

    { path: '**', redirectTo: 'login' }
];
