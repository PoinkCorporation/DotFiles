source /usr/share/cachyos-fish-config/cachyos-config.fish

# overwrite greeting
# potentially disabling fastfetch
#function fish_greeting
#    # smth smth
#end

function esp
    # Allow git to run inside the /opt/esp-idf folder
    git config --global --get-all safe.directory | grep -q '^/opt/esp-idf$'
    or git config --global --add safe.directory /opt/esp-idf

    # Source the fish-compatible export script
    source /opt/esp-idf/export.fish
end


# >>> ESP-IDF EIM PATH >>>
# Added by ESP-IDF extension so the EIM CLI can be launched directly.
if not contains -- "/home/rosetka/.espressif/eim_gui" $PATH
    set -gx PATH "/home/rosetka/.espressif/eim_gui" $PATH
end
# <<< ESP-IDF EIM PATH <<<
