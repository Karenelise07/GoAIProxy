let currentBot = '';
let currentBotAvatar = '';
const userAvatar = './user.jpg'; // 假设用户头像URL
const botConversations = {
  'Bot 1': [],
  'Bot 2': [],
  'Bot 3': []
};

function selectBot(botName, botAvatar, index) {
  currentBot = botName;
  currentBotAvatar = botAvatar;

  // 更新选中的机器人样式
  const botListItems = document.getElementById('botList').children;
  for (let i = 0; i < botListItems.length; i++) {
    botListItems[i].classList.remove('selected');
  }
  botListItems[index].classList.add('selected');

  // 更新当前机器人信息区域
  document.getElementById('currentBotName').textContent = botName;
  document.getElementById('currentBotAvatar').src = botAvatar;

  // 清空对话并加载当前选中的机器人对话历史
  const messagesContainer = document.getElementById('messages');
  messagesContainer.innerHTML = ''; // 清空当前对话
  const conversations = botConversations[currentBot] || [];
  conversations.forEach(message => {
    if (message.sender === 'user') {
      messagesContainer.innerHTML += `<div class="message user"><div class="text">${message.text}</div><img class="avatar" src="${userAvatar}" alt="User"></div>`;
    } else {
      messagesContainer.innerHTML += `<div class="message bot"><img class="avatar" src="${currentBotAvatar}" alt="${currentBot}"><div class="text">${message.text}</div></div>`;
    }
  });

  // 滚动到最新消息
  messagesContainer.scrollTop = messagesContainer.scrollHeight;
}


function sendMessage() {
  const input = document.getElementById('userInput');
  const message = input.value.trim();
  if (message && currentBot) {
    const userMessage = { sender: 'user', text: message };
    botConversations[currentBot].push(userMessage);
    
    // 显示用户消息
    const messagesContainer = document.getElementById('messages');
    messagesContainer.innerHTML += `<div class="message user"><div class="text">${message}</div><img class="avatar" src="${userAvatar}" alt="User"></div>`;
    
    // 发送消息到服务器并处理回复
    sendToBotServer(message);

    input.value = ''; // 清空输入框
  }
}


function sendToBotServer(userMessage) {
  fetch('http://172.21.44.125:8081/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ user_message: userMessage }),
  })
  .then(response => response.json())
  .then(data => {
    const botReply = { sender: 'bot', text: data.response };
    botConversations[currentBot].push(botReply);
    
    // 显示机器人回复
    const messagesContainer = document.getElementById('messages');
    messagesContainer.innerHTML += `<div class="message bot"><img class="avatar" src="${currentBotAvatar}" alt="${currentBot}"><div class="text">${botReply.text}</div></div>`;

    // 滚动到最新消息
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
  })
  .catch(error => {
    console.error('Error sending message to bot:', error);
  });
}



function init() {
  // 监听回车键发送消息
  document.getElementById('userInput').addEventListener('keypress', function(event) {
    if (event.key === 'Enter') {
      event.preventDefault(); // 阻止默认行为
      sendMessage();
    }
  });
}

// 在页面加载完成后执行init函数
window.onload = init;
