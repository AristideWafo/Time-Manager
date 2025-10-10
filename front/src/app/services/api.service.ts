import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class ApiService {

    constructor(private http: HttpClient) { }

    login(credentials: { username: string; password: string }) {
        return this.http.post<{ token: string }>(
            `/authentification`,
            credentials
        );
    }

    getUser(id: number) {
        return this.http.get(`/user/${id}`);
    }

    postUser(data: any) {
        return this.http.post(`/user/`, data)
    }

    updateUser(id: number, data: any) {
        return this.http.put(`/user/${id}`, data);
    }

    deleteUser(id: number) {
        return this.http.delete(`/user/${id}`);
    }

    getTeamWorkers(teamId: number) {
        return this.http.get(`/team/${teamId}/workers`);
    }

    postTeamWorkers(teamId: number, data: any) {
        return this.http.post(`/team/${teamId}`, data)
    }

    updateTeamWorkers(teamId: number, data: any) {
        return this.http.put(`/team/${teamId}/workers`, data);
    }

    deleteTeamWorkers(teamId: number) {
        return this.http.delete(`/team/${teamId}`);
    }

    updatePointage(id: number, data: boolean) {
        return this.http.put(`/pointage/${id}`, data)
    }
}