describe('First Test', () => {
  it('Visit the App', () => {
    cy.visit('http://localhost:3001/')

    cy.get('input[name="ID"]').type('1').should('have.value','1')
    .get('input[name="Task"]').type('Github Actions').should('have.value','Github Actions')
    cy.get('.button').click()


    cy.get('input[name="ID"]').type('2').should('have.value','2')
    .get('input[name="Task"]').type('Docker Image').should('have.value','Docker Image')
    cy.get('.button').click()

    cy.get('input[name="ID"]').type('3').should('have.value','3')
      .get('input[name="Task"]').type('Documentation').should('have.value','Documentation')
    cy.get('.button').click()

    // cy.get('input[type="checkbox"]').check()
    cy.contains('❌').click()

  })
})

