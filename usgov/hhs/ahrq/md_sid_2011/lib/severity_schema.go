package lib

var severitySchema = 
`Data Set Name: MD_SIDC_2011_SEVERITY           
Number of Observations: 721317
Total Record Length: 76
Total Number of Variables: 30


Columns   Description:
=======   ============
 1-  3    Database name
 5-  8    Discharge year of data
10- 25    File name
27- 29    Variable number
31- 56    Variable name
58- 61    Starting column of variable in ASCII file
63- 66    Ending column of variable in ASCII file
    68    Non-zero number of digits after decimal point for numeric variable
70- 73    Variable type (Num=numeric; Char=character)
75-174    Variable label


MD  2011 SEVERITY           1 KEY                           1   18   Num  HCUP record identifier
MD  2011 SEVERITY           2 CM_AIDS                      19   20   Num  AHRQ comorbidity measure: Acquired immune deficiency syndrome
MD  2011 SEVERITY           3 CM_ALCOHOL                   21   22   Num  AHRQ comorbidity measure: Alcohol abuse
MD  2011 SEVERITY           4 CM_ANEMDEF                   23   24   Num  AHRQ comorbidity measure: Deficiency anemias
MD  2011 SEVERITY           5 CM_ARTH                      25   26   Num  AHRQ comorbidity measure: Rheumatoid arthritis/collagen vascular diseases
MD  2011 SEVERITY           6 CM_BLDLOSS                   27   28   Num  AHRQ comorbidity measure: Chronic blood loss anemia
MD  2011 SEVERITY           7 CM_CHF                       29   30   Num  AHRQ comorbidity measure: Congestive heart failure
MD  2011 SEVERITY           8 CM_CHRNLUNG                  31   32   Num  AHRQ comorbidity measure: Chronic pulmonary disease
MD  2011 SEVERITY           9 CM_COAG                      33   34   Num  AHRQ comorbidity measure: Coagulopathy
MD  2011 SEVERITY          10 CM_DEPRESS                   35   36   Num  AHRQ comorbidity measure: Depression
MD  2011 SEVERITY          11 CM_DM                        37   38   Num  AHRQ comorbidity measure: Diabetes, uncomplicated
MD  2011 SEVERITY          12 CM_DMCX                      39   40   Num  AHRQ comorbidity measure: Diabetes with chronic complications
MD  2011 SEVERITY          13 CM_DRUG                      41   42   Num  AHRQ comorbidity measure: Drug abuse
MD  2011 SEVERITY          14 CM_HTN_C                     43   44   Num  AHRQ comorbidity measure: Hypertension (combine uncomplicated and complicated)
MD  2011 SEVERITY          15 CM_HYPOTHY                   45   46   Num  AHRQ comorbidity measure: Hypothyroidism
MD  2011 SEVERITY          16 CM_LIVER                     47   48   Num  AHRQ comorbidity measure: Liver disease
MD  2011 SEVERITY          17 CM_LYMPH                     49   50   Num  AHRQ comorbidity measure: Lymphoma
MD  2011 SEVERITY          18 CM_LYTES                     51   52   Num  AHRQ comorbidity measure: Fluid and electrolyte disorders
MD  2011 SEVERITY          19 CM_METS                      53   54   Num  AHRQ comorbidity measure: Metastatic cancer
MD  2011 SEVERITY          20 CM_NEURO                     55   56   Num  AHRQ comorbidity measure: Other neurological disorders
MD  2011 SEVERITY          21 CM_OBESE                     57   58   Num  AHRQ comorbidity measure: Obesity
MD  2011 SEVERITY          22 CM_PARA                      59   60   Num  AHRQ comorbidity measure: Paralysis
MD  2011 SEVERITY          23 CM_PERIVASC                  61   62   Num  AHRQ comorbidity measure: Peripheral vascular disorders
MD  2011 SEVERITY          24 CM_PSYCH                     63   64   Num  AHRQ comorbidity measure: Psychoses
MD  2011 SEVERITY          25 CM_PULMCIRC                  65   66   Num  AHRQ comorbidity measure: Pulmonary circulation disorders
MD  2011 SEVERITY          26 CM_RENLFAIL                  67   68   Num  AHRQ comorbidity measure: Renal failure
MD  2011 SEVERITY          27 CM_TUMOR                     69   70   Num  AHRQ comorbidity measure: Solid tumor without metastasis
MD  2011 SEVERITY          28 CM_ULCER                     71   72   Num  AHRQ comorbidity measure: Peptic ulcer disease excluding bleeding
MD  2011 SEVERITY          29 CM_VALVE                     73   74   Num  AHRQ comorbidity measure: Valvular disease
MD  2011 SEVERITY          30 CM_WGHTLOSS                  75   76   Num  AHRQ comorbidity measure: Weight loss
`