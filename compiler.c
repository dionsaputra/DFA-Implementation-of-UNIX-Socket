#include <stdio.h>

/******** KAMUS TOKEN *************/
char token[50][50];
	token[1] = "program";	token[2] = "var";		token[3] = "integer";	token[4] = "real";
	token[5] = "char";		token[6] = "array"; 	token[7] = "of";		token[8] = "begin";
	token[9] = "end";		token[10] = ":";		token[11] = ";";		token[12] = ",";
	token[13] = ".";		token[14] = ":=";		token[15] = "+";		token[16] = "-";
	token[17] = "*";		token[18] = "div";		token[19] = "mod";		token[20] = "if";
	token[21] = "then";		token[22] = "else";		token[23] = "while";	token[24] = "do";
	token[25] = "for";		token[26] = "(";		token[27] = "to";		token[28] = "downto";
	token[29] = ")";		token[30] = "step";		token[31] = "repeat";	token[32] = "until";
	token[33] = "input";	token[34] = "output";	token[35] = "and";		token[36] = "or";
	token[37] = "<";		token[38] = ">";		token[39] = "<=";		token[40] = ">=";
	token[41] = "=";		token[42] = "<>";		token[43] = "[";		token[44] = "]";

int length(char* str){
	int i=0;
	while ((int)str[i] != 0) i++;
	return i;
}

boolean equal(char* str1, char* str2){
	if (length(str1) != length(str2)){
		return false;
	} else {
		int i=0;
		while (i<len(str1) && str1[i]==str2[i]) i++;
		return str1[i] == str2[i];
	}
}

int get_id(char* statement){
	int i=1;
	while (i<45){
		
		i++;
	}
}

int main(){ 
	char statement[100][100];
	FILE *f;
	f = fopen("syntax.txt","r");
	int i=0;
	do{
		scanf("%s",statement[i]);
		i++;
	} while !(feof(f));
	fclose(f);	
	for (int j=0; j<=i; j++){
		printf("%s\n",statement[j]);
	}
}