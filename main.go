package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type UDPBroadcaster struct {
	ipEntry      *widget.Entry
	portEntry    *widget.Entry
	messageEntry *widget.Entry
	timeoutEntry *widget.Entry

	logLabel     *widget.Label
	logContainer *container.Scroll

	logText string

	sendBtn     *widget.Button
	statusLabel *widget.Label
}

func NewUDPBroadcaster() *UDPBroadcaster {
	ub := &UDPBroadcaster{}

	// 输入
	ub.ipEntry = widget.NewEntry()
	ub.ipEntry.SetText("255.255.255.255")

	ub.portEntry = widget.NewEntry()
	ub.portEntry.SetText("48899")

	ub.messageEntry = widget.NewEntry()
	ub.messageEntry.SetText("www.usr.cn")

	ub.timeoutEntry = widget.NewEntry()
	ub.timeoutEntry.SetText("5")

	// ✅ 日志（核心改动）
	ub.logLabel = widget.NewLabel("")
	ub.logLabel.Wrapping = fyne.TextWrapWord

	ub.logContainer = container.NewScroll(ub.logLabel)
	ub.logContainer.SetMinSize(fyne.NewSize(700, 300))

	ub.statusLabel = widget.NewLabel("就绪")

	ub.sendBtn = widget.NewButton("发送广播", ub.sendBroadcast)

	return ub
}

func (ub *UDPBroadcaster) updateUI(f func()) {
	fyne.Do(f)
}

func (ub *UDPBroadcaster) addLog(msg string) {
	line := time.Now().Format("15:04:05 ") + msg

	ub.updateUI(func() {
		ub.logText += line + "\n"
		ub.logLabel.SetText(ub.logText)
		ub.logContainer.ScrollToBottom()
	})
}

func (ub *UDPBroadcaster) sendBroadcast() {

	broadcastIP := ub.ipEntry.Text
	port := ub.portEntry.Text
	message := ub.messageEntry.Text
	timeout := ub.timeoutEntry.Text

	ub.sendBtn.Disable()
	ub.statusLabel.SetText("发送中...")

	go func() {

		defer func() {
			ub.updateUI(func() {
				ub.sendBtn.Enable()
				ub.statusLabel.SetText("就绪")
			})
		}()

		portInt, _ := strconv.Atoi(port)
		timeoutInt, _ := strconv.Atoi(timeout)

		localAddr, _ := net.ResolveUDPAddr("udp", ":0")
		conn, err := net.ListenUDP("udp", localAddr)
		if err != nil {
			ub.addLog(fmt.Sprintf("❌ UDP创建失败: %v", err))
			return
		}
		defer conn.Close()

		udpAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", broadcastIP, portInt))

		_, err = conn.WriteToUDP([]byte(message), udpAddr)
		if err != nil {
			ub.addLog(fmt.Sprintf("❌ 发送失败: %v", err))
			return
		}

		ub.addLog(fmt.Sprintf("📡 广播发送 → %s:%d", broadcastIP, portInt))

		conn.SetReadDeadline(time.Now().Add(time.Duration(timeoutInt) * time.Second))

		buffer := make([]byte, 2048)
		count := 0

		for {
			n, addr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				break
			}

			count++
			resp := strings.TrimSpace(string(buffer[:n]))
			ub.addLog(fmt.Sprintf("📨 [%d] %s → %s", count, addr.String(), resp))
		}

		if count == 0 {
			ub.addLog("⏰ 未发现设备")
		} else {
			ub.addLog(fmt.Sprintf("✅ 完成，共发现 %d 个设备", count))
		}
	}()
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("UDP设备发现工具")
	myWindow.Resize(fyne.NewSize(750, 520))

	b := NewUDPBroadcaster()

	form := widget.NewForm(
		widget.NewFormItem("广播地址", b.ipEntry),
		widget.NewFormItem("端口", b.portEntry),
		widget.NewFormItem("消息", b.messageEntry),
		widget.NewFormItem("超时(秒)", b.timeoutEntry),
	)

	content := container.NewBorder(
		container.NewVBox(
			form,
			container.NewHBox(
				b.sendBtn,
				widget.NewButton("清空日志", func() {
					b.logText = ""
					b.logLabel.SetText("")
				}),
			),
		),
		container.NewHBox(b.statusLabel),
		nil,
		nil,
		container.NewBorder(
			widget.NewLabel("📋 日志"),
			nil,
			nil,
			nil,
			b.logContainer,
		),
	)

	myWindow.SetContent(content)

	b.addLog("工具已启动")
	b.addLog("点击“发送广播”开始扫描设备")

	myWindow.ShowAndRun()
}
