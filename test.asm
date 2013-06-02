*        =   $10
EINS     DS  1
;M_EINS   DS  -1
MAX      =   100
ZAEHLER  DS  ; Zaehlervariable

*        =   128
; Hauptprogramm fuer eine Zaehlschleife
START    LDV EINS     ; Initialisieren ZAEHLER mit 1
         STV ZAEHLER
SCHLEIFE LDV ZAEHLER  ; erhoehe ZAEHLER um 1
         ADD EINS
         STV ZAEHLER
         LDC MAX      ; falls MAX erreicht ist
         EQL ZAEHLER
         JMN FERTIG   ; dann beende die Schleife
         JMP SCHLEIFE ; sonst naechster Schleifendurchlauf
FERTIG   HALT


