//Description : Le module "timeout" a pour fonction de pr�venir l'utilisateur d'un �tat d'attente du contr�leur d'un certains nombre de donn�es 
//				suppos�es �tre recues. En effet, dans le cas de l'ecriture de n donn�es en m�moire le module timeout dispose d'un minuteur qui d�pass�
//				un d�lai, averti � l'aide d'un signal de contr�le que le nombre de donn�es re�ues est inf�rieure au nombre attendu.  



module a_timeout(
              clk_ref,
              rst_n,
              start_i,
              stop_i,
              r_erreur_timeout_o
              );

input clk_ref, rst_n, start_i, stop_i;
output r_erreur_timeout_o;

wire clk_ref, rst_n, start_i, stop_i;
reg r_erreur_timeout_o;

// ## DECLARATION VARIABLE INTERNE ###
reg[19:0] r_cpt_timeout;
reg r_demarre_timeout;

always @(posedge clk_ref or negedge rst_n)
  if(!rst_n) 
         begin
            r_erreur_timeout_o <= 1'b0;
            r_cpt_timeout <= 20'b0;
            r_demarre_timeout <= 1'b0;
         end   
  else if(start_i && !stop_i && !r_erreur_timeout_o)
         begin
            r_erreur_timeout_o <= 1'b0;
            r_cpt_timeout <= 20'b0;
            r_demarre_timeout <= 1'b1;
         end
  else if(stop_i && !r_erreur_timeout_o)
         begin
            r_erreur_timeout_o <= 1'b0;
            r_cpt_timeout <= 20'b0;
            r_demarre_timeout <= 1'b0;
         end
  else if(r_demarre_timeout && !r_erreur_timeout_o)
         begin
            r_erreur_timeout_o <= (r_cpt_timeout == 20'd500000) ? 1'b1 : 1'b0;
            r_cpt_timeout <= r_cpt_timeout + 1'b1;
            r_demarre_timeout <= 1'b1;
         end
  else if(r_demarre_timeout && r_erreur_timeout_o)
         begin
            r_erreur_timeout_o <= 1'b1;
            r_cpt_timeout <= r_cpt_timeout;
            r_demarre_timeout <= 1'b1;
         end
  else 
         begin
            r_erreur_timeout_o <= 1'b0;
            r_cpt_timeout <= 20'b0;
            r_demarre_timeout <= 1'b0;
         end         

endmodule
