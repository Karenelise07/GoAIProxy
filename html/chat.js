let currentBot = '';
let currentBotAvatar = '';
let accumulatedData = ''; // 用于累积接收到的数据
let currentBotMessageContainer = null; // 跟踪当前机器人消息的容器

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
    // 更新对话历史
    botConversations[currentBot].push(userMessage);
    
    // 显示用户消息
    const messagesContainer = document.getElementById('messages');
    appendMessage(messagesContainer, 'user', message, userAvatar);
    
    // 发送消息到服务器并处理回复
    sendToBotServer(message);

    input.value = ''; // 清空输入框
  }
}

function sendToBotServer(userMessage) {

  // 获取当前机器人的对话历史
  const currentConversations = botConversations[currentBot] || [];

  // 将对话历史转换为后端服务所需的格式
  const messages = currentConversations.map(message => ({
    content: message.text,
    role: message.sender === 'user' ? "user" : "system" // 使用 "system" 来标识机器人消息
  }));

  const requestData = {
    model: "qwen1.5-chat",
    messages: messages
  };

  fetch('http://localhost:8083/v1/chat/completions', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(requestData),
  }).then(response => {
    console.log('Received response from server');
    const reader = response.body.getReader();
    const decoder = new TextDecoder();
    let accumulatedContent = ''; // 用于累积对话内容

    function read() {
      reader.read().then(({done, value}) => {
        if (done) {
          console.log('Stream completed');
          return;
        }
        const chunk = decoder.decode(value, {stream: true});
        accumulatedContent += chunk; // 累积从流中读取的内容
        processChunk(chunk); // 处理接收到的每个数据块
        read(); // 继续读取下一个数据块
      }).catch(error => {
        console.error('Error reading from stream:', error);
      });
    }
    read();
  }).catch(error => {
    console.error('Error sending message to bot:', error);
  });
}


function processChunk(chunk) {
  accumulatedData += chunk; // 累积接收到的数据

  let start = 0; // 初始化查找起点
  let braceCounter = 0; // 大括号计数器
  let inObject = false; // 是否开始记录一个对象

  for (let i = 0; i < accumulatedData.length; i++) {
    if (accumulatedData[i] === '{') {
      braceCounter++;
      if (!inObject) {
        start = i; // 记录对象开始的位置
        inObject = true;
      }
    } else if (accumulatedData[i] === '}') {
      braceCounter--;
    }

    // 当计数器回到零时，我们找到了一个完整的JSON对象
    if (braceCounter === 0 && inObject) {
      // 我们找到了一个完整的JSON对象
      let jsonStr = accumulatedData.substring(start, i + 1); // 提取JSON字符串
      try {
        console.log("Trying to parse JSON:", jsonStr); // 打印出尝试解析的字符串
        let jsonObj = JSON.parse(jsonStr);
        // 处理解析出的JSON对象
        if (jsonObj.choices && jsonObj.choices.length > 0) {
          const choice = jsonObj.choices[0];
          if (choice.delta && choice.delta.content) {
            // 假设我们有一个函数来逐字展示消息
            appendMessage(document.getElementById('messages'), 'system', choice.delta.content, currentBotAvatar);

            //appendMessage(choice.delta.content, document.getElementById('messages'));
          }
          // 检查是否是对话结束的信号
          if (choice.finish_reason === 'stop') {
            // 如果是结束信号，可以在这里执行一些清理工作或者标记对话结束
            console.log('Conversation ended.');
          }
        }
      } catch (error) {
        console.error('Error parsing JSON chunk:', error);
      }
      // 准备查找下一个JSON对象
      inObject = false; // 重置状态
      accumulatedData = accumulatedData.substring(i + 1); // 移除已处理的部分
      i = -1; // 重置索引，因为我们修改了accumulatedData
    }
  }
}

function appendMessage(container, sender, text, avatar) {
  let messageElement = document.createElement('div');
  messageElement.classList.add('message', sender); // 使用 sender 作为类名，以便于 CSS 样式定制
  
  let avatarElement = document.createElement('img');
  avatarElement.classList.add('avatar');
  avatarElement.src = avatar;
  avatarElement.alt = sender === 'user' ? 'User' : 'System'; // 根据 sender 调整 alt 文本
  
  let textElement = document.createElement('div');
  textElement.classList.add('text');
  
  // 根据 sender 调整消息元素的顺序
  if (sender === 'user') {
    messageElement.appendChild(textElement);
    messageElement.appendChild(avatarElement);
  } else { // 假设 sender 为 'system'
    messageElement.appendChild(avatarElement);
    messageElement.appendChild(textElement);
  }
  
  container.appendChild(messageElement);
  typeMessage(text, textElement); // 确保您有一个实现了逐字显示文本的 typeMessage 函数
}

function typeMessage(message, container, index = 0) {
  if (index < message.length) {
    let textElement = document.createElement('span'); // 创建一个新的span元素来容纳当前字符
    textElement.textContent = message.charAt(index);
    //textElement.classList.add('system-message'); // 假设你已经定义了相应的CSS类
    container.appendChild(textElement); // 将span元素添加到容器中

    setTimeout(() => typeMessage(message, container, index + 1), 50); // 递归调用以展示下一个字符
  }
}


function init() {
  // 默认选择Bot 1
  selectBot('Bot 1', './bot.jpg', 0);
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
