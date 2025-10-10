import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';



@Injectable({ providedIn: 'root' })
export class ApiService {
    private baseUrl = ''; // mettre URL de API quand on aura le backend

    constructor(private http: HttpClient) { }

    getUserProfile(id: number) {
        return this.http.get(`${this.baseUrl}/profile/${id}`);
    }

    postUserProfile(id: number, data: any) {
        return this.http.post(`${this.baseUrl}/user/`, data)
    }

    updateUserProfile(id: number, data: any) {
        return this.http.put(`${this.baseUrl}/profile/${id}`, data);
    }

    deleteUserProfile(id: number) {
        return this.http.delete(`${this.baseUrl}/profile/${id}`);
    }

    getTeamWorkers(teamId: number) {
        return this.http.get(`${this.baseUrl}/team/${teamId}/workers`);
    }

    postTeamWorkers(teamId: number, data: any) {
        return this.http.post(`${this.baseUrl}/team/${teamId}`, data)
    }

    updateTeamWorkers(teamId: number, data: any) {
        return this.http.put(`${this.baseUrl}/team/${teamId}/workers`, data);
    }

    deleteTeamWorkers(teamId: number) {
        return this.http.delete(`${this.baseUrl}/team/${teamId}`);
    }



}




