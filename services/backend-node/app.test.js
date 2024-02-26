const { helloWorld } = require('./app');

test('helloWorld function should return "Hello, World!"', () => {
   expect(helloWorld()).toBe('Hello, World!');
});
