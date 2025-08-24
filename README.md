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

---

Dúvidas ou sugestões? Consulte os comentários nas interfaces ou entre em contato com o time responsável.

Infra sibling module for small infra helpers used by tests.

This module is intentionally minimal; add packages here for shared test helpers or infra-level code.
