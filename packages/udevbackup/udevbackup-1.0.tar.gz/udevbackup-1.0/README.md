UdevBackup
==========

On Linux, detects when specified storage devices are connected, mounts them,
executes a script, umounts them and tells when it is done.

A config file defines storage devices and the scripts to run.

I wrote this script for a simple offline backup of my server. I just turn
the external USB disk on and wait for the message (using text to speech) before
turning it off again. UdevBackup double forks before running the script, so
there is no timeout problem with slow scripts.

installation:

    sudo pip3 install udevbackup --upgrade

you need to create to launch udevbackup when a new device (with a file system) is connected:

    echo 'ACTION=="add", ENV{DEVTYPE}=="partition", RUN+="/usr/local/bin/udevbackup"' | sudo tee /etc/udev/rules.d/
    udevadm control --reload-rules

then create a config file (with the common .ini syntax).
Create a "main" section for global options, and another section for each
target partition. The name is not important.

You can display all available options with the "help" command, but .


    udevbackup help

    [main]
    smtp_auth_user = SMTP user. Default to "".
    smtp_auth_password = SMTP password. Default to "".
    smtp_server = SMTP server. Default to "localhost".
    smtp_from_email = Recipient of the e-mail.  Default to "".
    smtp_to_email = E-mail address for the FROM: value. Default to "".
    use_speech = Use google speech for announcing successes and failures. Default to 0.
    use_stdout = Display messages on stdout. Default to 0.
    use_smtp = Send messages by email (with the whole content of stdout/stderr of your scripts). Default to 0.
    smtp_use_tls = Use TLS (smtps) for emails. Default to 0.
    smtp_use_starttls = Use STARTTLS for emails. Default to 0.
    smtp_smtp_port = The SMTP port. Default to 25.

    [example]
    fs_uuid = UUID of the used file system. Check /dev/disk/by-uuid/ before and after having connected your disk.
    command = Command to call for running the script (whose name is passed as first argument). Default to "bash".
    script = Content of the script to execute when the disk is mounted. Current working dir is the mounted directory.
    stdout = Write stdout to this filename.
    stderr = Write stderr to this filename.
    mount_options = Extra mount options. Default to "".
    user = User used for running the script and mounting the disk.Default to "current user".
    pre_script = Script to run before mounting the disk. The disk will not be mounted if this script does not returns 0. Default to "".
    post_script = Script to run after the disk umount. Only run if the disk was mounted. Default to "".

Here is a complete example:

    cat /etc/udevbackup/example.ini
    [main]
    smtp_auth_user = user
    smtp_auth_password = s3cr3tP@ssw0rd
    smtp_server = localhost
    use_speech = 1
    use_stdout = 0
    use_smtp = 1

    [example]
    fs_uuid = 58EE-7CAE
    script = mkdir -p ./data
        rsync -av /data/to_backup/ ./data/


