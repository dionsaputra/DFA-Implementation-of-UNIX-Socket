#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int main(){
	//KAMUS 
	int num_state, num_input, num_route, num_fstate;
	int num_uinput=0;
	FILE *fp;
	char uinput[100][100];

	while(scanf("%s",uinput[num_uinput++]), strcmp(uinput[num_uinput-1],"end")!=0);
	
	//BACA STATE
	fp = freopen("state.txt","r",stdin);
	char state[100][100];
	scanf("%d", &num_state);
	for(int i=0;i<num_state;i++) scanf("%s", state[i]);

	//BACA INPUT
	fp = freopen("input.txt","r",stdin);
	char input[100][100];
	scanf("%d",&num_input);
	for(int i=0;i<num_input;i++) scanf("%s", input[i]);

	//BACA STATE AWAL
	fp = freopen("initState.txt","r",stdin);
	char cur[100], ac[100], inp[100];
	int cur_i = 0;
	scanf("%s", cur);

	//BACA FINAL STATE
	fp = freopen("finalState.txt","r",stdin);
	char fstate[100][100];
	scanf("%d", &num_fstate);
	for(int i=0;i<num_fstate;i++) scanf("%s", fstate[i]);
	
	//BACA LINTASAN
	fp = freopen("lintasan.txt","r",stdin);
	int route[100][100];
	scanf("%d",&num_route);

	//INISIALISASI MATRIKS DENGAN -1
	for(int i=0; i<num_state; i++){
		for(int j=0; j<num_input; j++){
			route[i][j]= -1;
		}
	}

	for(int i=0; i<num_route; i++){
		char u[100], e[100], v[100];
		scanf("%s %s %s", u, e, v);
		
		int idx_u = 0;
		while(strcmp(state[idx_u],u)!=0){
			idx_u++;
		}
		
		int idx_e = 0;
		while(strcmp(input[idx_e],e)!=0){
			idx_e++;
		}
		
		int idx_v = 0;
		while(strcmp(state[idx_v],v)!=0){
			idx_v++;
		}
		
		route[idx_u][idx_e] = idx_v;
	}

	fclose(fp);

	//VALIDASI APAKAH INPUT USER VALID?
	int valid_input1 = 1;
	for (int i=0; i<num_uinput-1; i++){
		int valid_input2 = 0;
		for (int j=0; j<num_input; j++){
			if (strcmp(uinput[i],input[j]) == 0){
				valid_input2 = 1;
			}
		}
		valid_input1 = (valid_input1 && valid_input2);
	}
	if (valid_input1){
		int cek_valid1 = 1;
		int cek_valid2 = 1;
		int arr_cur_i[100];
		for(int i=0; i<num_uinput-1; i++){
			int idx_ii = 0;
		
			while(strcmp(input[idx_ii],uinput[i])!=0 && idx_ii<num_input-1){
				idx_ii++;
			} 
			
			if(route[cur_i][idx_ii] == -1){
				cek_valid2 = 0;
			} else {
				arr_cur_i[i] = route[cur_i][idx_ii];
				cur_i = route[cur_i][idx_ii];
			}
			cek_valid1 = (cek_valid1 && cek_valid2);
		} 
		
		
		int cek_valid3 = 0;
		for (int i=0; i<num_fstate && cek_valid3==0; i++){
			if (strcmp(state[cur_i],fstate[i])==0){
				cek_valid3 = 1;
			}
		}

		if (cek_valid1 == 1 && cek_valid3 == 1){
			printf("\nAccepted\n");

			printf("\nLintasan : \n");
			printf("State_0 + %s -> %s\n",uinput[0], state[arr_cur_i[0]]);
			for (int i=1; i<num_uinput-1; i++){
				printf("%s + %s -> %s\n",state[arr_cur_i[i-1]], uinput[i], state[arr_cur_i[i]]);
			}
		} else {
			printf("\nRejected\n");
		}
	} else {
		printf("Terdapat input user yang tidak ada pada daftar input DFA\n");
	}
	
	//APP.EXE diberi waktu jeda untuk menampilkan hasil
	for (int i=0; i<1e9; i++){}
	return 0;
}
