from typing import Any, Callable, List, Union

from pyrogram import ChatAction
from pyrogram.api.types import BotCommand, PeerUser
from pyrogram.client.filters.filter import Filter

from .containers import InlineResultContainer
from .interactionclient import InteractionClient
from .response import Response


class BotIntegrationClient(InteractionClient):
    bot_under_test: Union[int, str] = ...
    max_wait_response: float = ...

    min_wait_consecutive: float = ...
    global_action_delay: float = ...
    peer: PeerUser = ...
    peer_id: int = ...
    command_list: List[BotCommand] = ...

    _last_response: Response = ...

    # do you know a a way to automate this?

    def __init__(
            self,
            bot_under_test: Union[int, str],
            session_name: str = None,
            api_id: int = None,
            api_hash: str = None,
            phone_number: str = None,
            max_wait_response: int = None,
            min_wait_consecutive: float = None,
            raise_no_response: bool = None,
            global_action_delay: int = None,
            workdir: str = '.',
            config_file="./config.ini",
            **kwargs: Any
    ) -> Any:
        super().__init__(session_name, api_id, api_hash, phone_number, workdir=workdir,
                         config_file=config_file)
        ...

    def start(self, debug: bool = ...):
        ...

    def send_audio_await(
            self,
            audio: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            caption: str = ...,
            parse_mode: str = ...,
            duration: int = ...,
            performer: str = ...,
            title: str = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response:
        ...

    def send_chat_action_await(
            self,
            action: ChatAction or str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response:
        ...

    def send_contact_await(
            self,
            chat_id: int or str,
            phone_number: str,
            first_name: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            last_name: str = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            **kwargs
    ) -> Response:
        ...

    def send_document_await(
            self,
            document: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            caption: str = ...,
            parse_mode: str = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response: ...

    def send_location_await(
            self,
            latitude: float,
            longitude: float,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            **kwargs
    ) -> Response: ...

    def send_media_group_await(
            self,
            media: list,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            **kwargs
    ) -> Response: ...

    def send_message_await(
            self,
            text,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            **kwargs
    ) -> Response:
        ...

    def send_command_await(
            self,
            command: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            **kwargs
    ) -> Response:
        ...

    def send_photo_await(
            self,
            photo: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            caption: str = ...,
            parse_mode: str = ...,
            ttl_seconds: int = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response:
        ...

    def send_sticker_await(
            self,
            sticker: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response:
        ...

    def send_venue_await(
            self,
            latitude: float,
            longitude: float,
            title: str,
            address: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            foursquare_id: str = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            **kwargs
    ) -> Response:
        ...

    def send_video_await(
            self,
            video: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            caption: str = ...,
            parse_mode: str = ...,
            duration: int = ...,
            width: int = ...,
            height: int = ...,
            thumb: str = ...,
            supports_streaming: bool = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response:
        ...

    def send_video_note_await(
            self,
            video_note: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            duration: int = ...,
            length: int = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response:
        ...

    def send_voice_await(
            self,
            voice: str,
            filters: Filter = ...,
            num_expected: int = ...,
            raise_: bool = ...,
            caption: str = ...,
            parse_mode: str = ...,
            duration: int = ...,
            disable_notification: bool = ...,
            reply_to_message_id: int = ...,
            progress: Callable = ...,
            **kwargs
    ) -> Response:
        ...

    def ping(self, override_messages: List[str] = None) -> Response:
        """
        Send messages to a bot to determine whether it is online.

        Specify a list of ``override_messages`` that should be sent to the bot, defaults to /start.

        Args:
            override_messages: List of messages to be sent, defaults to /start.

        Returns:
            Response
        """
        ...

    def clear_chat(self) -> None:
        ...

    def _make_awaitable_method(self, name, method) -> None:
        ...

    def _get_command_list(self) -> List[BotCommand]:
        ...

    def get_default_filters(self, user_filters: Filter = ...) -> Filter:
        ...

    def get_inline_results(
            self,
            query: str,
            offset: str = '',
            latitude: int = None,
            longitude: int = None
    ) -> InlineResultContainer:
        ...
