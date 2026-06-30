// Injeção de dependência com interfaces
package main

import (
	"errors"
	"fmt"
)

// =====================================================
// INTERFACES / CONTRATOS
// =====================================================
//
// Aqui ficam os comportamentos que a camada de serviço precisa.
//
// O serviço NÃO quer saber se os dados vêm de:
// - memória
// - banco de dados
// - API externa
// - arquivo
//
// Ele só quer alguém que saiba buscar paciente por CPF.

type BuscadorDePaciente interface {
	BuscarNomePorCPF(cpf string) (string, bool)
}

// O serviço também não quer saber se o log vai para:
// - terminal
// - arquivo
// - Datadog
// - CloudWatch
//
// Ele só quer alguém que saiba registrar uma mensagem.

type Registrador interface {
	Registrar(mensagem string)
}

// =====================================================
// REPOSITORY / INFRAESTRUTURA DE DADOS
// =====================================================
//
// Essa parte representa uma implementação concreta de armazenamento.
//
// Em projeto real, poderia ser:
//
// type RepositorioPostgres struct { db *sql.DB }
//
// Mas aqui usamos map para simplificar.

type RepositorioMemoria struct {
	pacientes map[string]string
}

// Função construtora/factory.
// Ela cria e devolve uma struct concreta.

func NovoRepositorioMemoria() RepositorioMemoria {
	return RepositorioMemoria{
		pacientes: map[string]string{
			"111": "Luan",
			"222": "Ana",
			"333": "Joao",
		},
	}
}

// Esse método faz RepositorioMemoria satisfazer a interface BuscadorDePaciente.
//
// A struct NÃO escreve:
//
// implements BuscadorDePaciente
//
// Em Go, basta ter o método certo.

func (r RepositorioMemoria) BuscarNomePorCPF(cpf string) (string, bool) {
	nome, encontrado := r.pacientes[cpf]
	return nome, encontrado
}

// =====================================================
// LOGGER / INFRAESTRUTURA DE LOG
// =====================================================
//
// Outra implementação concreta.
// Aqui o log vai para o terminal.

type RegistradorTerminal struct{}

// Esse método faz RegistradorTerminal satisfazer a interface Registrador.

func (RegistradorTerminal) Registrar(mensagem string) {
	fmt.Println("log:", mensagem)
}

// =====================================================
// SERVICE LAYER / REGRA DE NEGÓCIO
// =====================================================
//
// Aqui fica a regra de negócio.
//
// O serviço depende de interfaces, não de structs concretas.
//
// Ele não sabe que existe RepositorioMemoria.
// Ele não sabe que existe RegistradorTerminal.
//
// Ele só conhece:
// - BuscadorDePaciente
// - Registrador

type ServicoDePacientes struct {
	buscador    BuscadorDePaciente
	registrador Registrador
}

// Construtor do serviço.
//
// Repare:
// aceita interfaces
// retorna struct concreta

func NovoServicoDePacientes(
	buscador BuscadorDePaciente,
	registrador Registrador,
) ServicoDePacientes {
	return ServicoDePacientes{
		buscador:    buscador,
		registrador: registrador,
	}
}

// Método com regra de negócio.
//
// Fluxo:
// 1. registra log
// 2. busca paciente por CPF
// 3. se não encontrar, retorna erro
// 4. se encontrar, monta saudação

func (s ServicoDePacientes) MontarSaudacao(cpf string) (string, error) {
	s.registrador.Registrar("buscando paciente com CPF " + cpf)

	nome, encontrado := s.buscador.BuscarNomePorCPF(cpf)
	if !encontrado {
		return "", errors.New("paciente nao encontrado")
	}

	return "Ola, " + nome + "!", nil
}

// =====================================================
// MAIN / COMPOSIÇÃO DA APLICAÇÃO
// =====================================================
//
// A main conecta tudo.
//
// Ela escolhe as implementações concretas:
// - RepositorioMemoria
// - RegistradorTerminal
//
// E injeta essas dependências no serviço.

func main() {
	// implementação concreta do repository
	repositorio := NovoRepositorioMemoria()

	// implementação concreta do logger
	registrador := RegistradorTerminal{}

	// service recebe interfaces, mas eu posso passar structs concretas
	// porque elas satisfazem essas interfaces
	servico := NovoServicoDePacientes(repositorio, registrador)

	// uso da regra de negócio
	saudacao, err := servico.MontarSaudacao("111")
	if err != nil {
		fmt.Println("erro:", err)
		return
	}

	fmt.Println(saudacao)
}
