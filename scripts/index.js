'use strict';

const textArea = document.getElementById('myTextarea')
const submitButton = document.getElementById('submit');
const greenButton = document.getElementById('green');
const moveButton = document.getElementById('movemovemove');

const getRequest = async (cmds) => {
  const url = `http://127.0.0.1:17000/?cmd=${cmds}`;
  await fetch(url, {
    mode: 'no-cors'
  })
    .then((res) => console.log(res))
    .catch((err) => console.log(err));
};

submitButton.onclick = async () => {
  const cmds = textArea.value.replace(/\n/g, ',');
  await getRequest(cmds);
};

greenButton.onclick = async () => {
  const cmds = 'green,bgrect 0.25 0.25 0.75 0.75,update';
  await getRequest(cmds);
};

moveButton.onclick = async () => {
  let i = 0;
  const cmds = 'white,figure 0.0 0.0,update';
  const cmdsToMove = 'move 0.1 0.1, update';
  await getRequest(cmds);

  const intervalToMove = setInterval(async () => {
    if (i === 10) {
      clearInterval(intervalToMove);
      i = 0;
    }
    await getRequest(cmdsToMove);
    i++;
  }, 1000)
};
