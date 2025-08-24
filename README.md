# Módulo infra

Este módulo provê abstrações de persistência para o projeto, usando GORM e interfaces para repositórios genéricos e específicos.

## Principais recursos
- Repositórios genéricos (`Repository[T]`) e específicos (ex: `AnexoRepo`)
- Interfaces públicas para fácil integração e testes
- Suporte a transações e métodos CRUD

## Como usar
```go
import (
    "github.com/brunojet/infra-go/pkg/repo"
    "gorm.io/gorm"
)

// Inicialize o GORM
var db *gorm.DB = ...

// Crie um repositório
repo := repo.NewRepository[MeuTipo](db)
err := repo.Create(ctx, &obj)
```

## Exemplo completo
```go
import (
    "context"
    "github.com/brunojet/infra-go/pkg/repo"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type MeuTipo struct {
    ID   int64
    Nome string
}

func main() {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    r := repo.NewRepository[MeuTipo](db)
    ctx := context.Background()
    obj := MeuTipo{Nome: "Exemplo"}
    _ = r.Create(ctx, &obj)
}
```

## Transações
```go
repo.WithTx(ctx, func(txRepo repo.Repository[MeuTipo]) error {
    // use txRepo normalmente
    return nil
})
```

## Testes
- Teste apenas os métodos essenciais (ex: Create)
- Use banco em memória para testes rápidos

## Boas práticas
- Sempre use as interfaces públicas para desacoplamento
- Trate erros de forma padronizada
- Consulte o README dos domínios para detalhes das entidades

## Testes de integração
Todos os métodos principais (Create, Get, Update, Delete, List) possuem exemplos de teste usando banco em memória:
```go
func TestRepository_CreateAndGet(t *testing.T) {
    db := setupDB(t)
    r := repo.NewRepository[MeuTipo](db)
    ctx := context.Background()
    obj := MeuTipo{Nome: "Teste"}
    err := r.Create(ctx, &obj)
    require.NoError(t, err)
    got, err := r.GetByID(ctx, obj.ID)
    require.NoError(t, err)
    require.Equal(t, obj.Nome, got.Nome)
}
```

---

Dúvidas ou sugestões? Consulte os comentários nas interfaces ou entre em contato com o time responsável.

Infra sibling module for small infra helpers used by tests.

This module is intentionally minimal; add packages here for shared test helpers or infra-level code.
