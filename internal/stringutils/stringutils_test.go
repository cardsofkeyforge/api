package stringutils

import "testing"

func TestEspecialTitle(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{"CotA", args{"Toma essa, Sabichão"}},
		{"CotA", args{"Irmãos de Batalha"}},
		{"CotA", args{"Queime o Estoque"}},
		{"CotA", args{"Desafio dos Campeões"}},
		{"CotA", args{"Ocaso do Covarde"}},
		{"CotA", args{"Lomir Punho de Fogo"}},
		{"CotA", args{"Um Jogo Justo"}},
		{"CotA", args{"Ergam-se!"}},
		{"CotA", args{"Chave para Dis"}},
		{"CotA", args{"Mestre de 1"}},
		{"CotA", args{"O Terror"}},
		{"CotA", args{"Explosão de PEM"}},
		{"CotA", args{"Æmber Irradiado"}},
		{"CotA", args{"Não se Alie a Marcianos"}},
		{"CotA", args{"Apoio da Nave-Mãe"}},
		{"CotA", args{"Minas à Distância"}},
		{"CotA", args{"Olho-da-Bala"}},
		{"AoA", args{"1-2 Soco"}},
		{"AoA", args{"À Luta"}},
		{"AoA", args{"Og, a Mestra Forjadora"}},
		{"AoA", args{"Diretor do Z.Y.X."}},
		{"AoA", args{"[CENSURADO]"}},
		{"AoA", args{"Matazord v. 9001"}},
		{"AoA", args{"Proclamação 346E"}},
		{"AoA", args{"Proteja os Fracos"}},
		{"AoA", args{"Uma Vida por outra Vida"}},
		{"AoA", args{"A Gripe Comum"}},
		{"WC", args{"O Chão é Lava"}},
		{"WC", args{"Ás dos Buracos de Minhoca"}},
		{"WC", args{"Unidade V.Æ.L.A."}},
		{"WC", args{"Dr. Milli"}},
		{"WC", args{"Ataque em Falange"}},
		{"WC", args{"A. Vinda"}},
		{"WC", args{"J. Vinda"}},
		{"WC", args{"Caça ou Caçador?"}},
		{"WC", args{"Oficial de Com. Kirby"}},
		{"WC", args{"CALV-1N"}},
		{"WC", args{"OIC Taber"}},
		{"WC", args{"Livro de IeQ"}},
		{"MM", args{"Q-Mecas"}},
		{"MM", args{"Dinosculpe-me"}},
		{"MM", args{"ANT1-10NY"}},
		{"DT", args{"ESLR Australis"}},
		{"DT", args{"EDAI “Edie” 4x4"}},
		{"DT", args{"EGS Luminoso"}},
		{"DT", args{"5C077"}},
		{"DT", args{"L30-P4RD0"}},
		{"DT", args{"Módulo 5UB-M3R50"}},
		{"DT", args{"T3R-35A"}},
		{"Rariry", args{"Rare"}},
		{"Rariry", args{"FIXED"}},
		{"Rariry", args{"Variant"}},
		{"Rariry", args{"Evil Twin"}},
		{"Rariry", args{"The Tide"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EspecialTitle(tt.args.value); got != tt.args.value {
				t.Errorf("EspecialTitle() = %v, want %v", got, tt.args.value)
			}
		})
	}
}
