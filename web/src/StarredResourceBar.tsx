import React from "react"
import { useHistory } from "react-router"
import styled from "styled-components"
import { ReactComponent as StarSvg } from "./assets/svg/star.svg"
import { usePathBuilder } from "./PathBuilder"
import { ClassNameFromResourceStatus } from "./ResourceStatus"
import { useSidebarPin } from "./SidebarPin"
import { combinedStatus } from "./status"
import {
  AnimDuration,
  Color,
  ColorAlpha,
  ColorRGBA,
  Font,
  FontSize,
  Glow,
  mixinResetButtonStyle,
  SizeUnit,
} from "./style-helpers"
import TiltTooltip from "./Tooltip"
import { ResourceStatus } from "./types"

export const StarredResourceLabel = styled.div`
  max-width: ${SizeUnit(4.5)};
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  display: inline-block;

  font-size: ${FontSize.small};
  font-family: ${Font.monospace};

  user-select: none;
`
const ResourceButton = styled.button`
  ${mixinResetButtonStyle};
  color: inherit;
  display: flex;
`
const StarIcon = styled(StarSvg)`
  height: ${SizeUnit(0.5)};
  width: ${SizeUnit(0.5)};
`
export const StarButton = styled.button`
  ${mixinResetButtonStyle};
  ${StarIcon} {
    fill: ${Color.grayLight};
  }
  &:hover {
    ${StarIcon} {
      fill: ${Color.grayLightest};
    }
  }
`
const StarredResourceRoot = styled.div`
  border-width: 1px;
  border-style: solid;
  border-radius: ${SizeUnit(0.125)};
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  background-color: ${Color.gray};
  padding-top: ${SizeUnit(0.125)};
  padding-bottom: ${SizeUnit(0.125)};

  &:hover {
    background-color: ${ColorRGBA(Color.gray, ColorAlpha.translucent)};
  }

  &.isSelected {
    background-color: ${Color.white};
    color: ${Color.gray};
  }

  &.isWarning {
    color: ${Color.yellow};
    border-color: ${ColorRGBA(Color.yellow, ColorAlpha.translucent)};
  }
  &.isHealthy {
    color: ${Color.green};
    border-color: ${ColorRGBA(Color.green, ColorAlpha.translucent)};
  }
  &.isUnhealthy {
    color: ${Color.red};
    border-color: ${ColorRGBA(Color.red, ColorAlpha.translucent)};
  }
  &.isBuilding {
    color: ${ColorRGBA(Color.white, ColorAlpha.translucent)};
  }
  .isSelected &.isBuilding {
    color: ${ColorRGBA(Color.gray, ColorAlpha.translucent)};
  }
  &.isPending {
    color: ${ColorRGBA(Color.white, ColorAlpha.translucent)};
    animation: ${Glow.white} 2s linear infinite;
  }
  .isSelected &.isPending {
    color: ${ColorRGBA(Color.gray, ColorAlpha.translucent)};
    animation: ${Glow.dark} 2s linear infinite;
  }
  &.isNone {
    color: ${Color.grayLighter};
    transition: border-color ${AnimDuration.default} linear;
  }

  // implement margins as padding on child buttons, to ensure the buttons consume the
  // whole bounding box
  ${StarButton} {
    margin-left: ${SizeUnit(0.25)};
    padding-right: ${SizeUnit(0.25)};
  }
  ${ResourceButton} {
    padding-left: ${SizeUnit(0.25)};
  }
`
const StarredResourceBarRoot = styled.div`
  padding-left: ${SizeUnit(0.5)};
  padding-right: ${SizeUnit(0.5)};
  padding-top: ${SizeUnit(0.25)};
  padding-bottom: ${SizeUnit(0.25)};
  margin-bottom: ${SizeUnit(0.25)};
  background-color: ${Color.grayDarker};

  ${StarredResourceRoot} {
    margin-right: ${SizeUnit(0.25)};
  }
`

export type ResourceNameAndStatus = {
  name: string
  status: ResourceStatus
}
export type StarredResourceBarProps = {
  selectedResource?: string
  resources: ResourceNameAndStatus[]
  unstar: (name: string) => void
}

export function StarredResource(props: {
  resource: ResourceNameAndStatus
  unstar: (name: string) => void
  isSelected: boolean
}) {
  const pb = usePathBuilder()
  const href = pb.encpath`/r/${props.resource.name}/overview`
  const history = useHistory()
  const onClick = (e: any) => {
    props.unstar(props.resource.name)
    e.preventDefault()
    e.stopPropagation()
  }

  let classes = [ClassNameFromResourceStatus(props.resource.status)]
  if (props.isSelected) {
    classes.push("isSelected")
  }

  return (
    <TiltTooltip title={props.resource.name}>
      <StarredResourceRoot className={classes.join(" ")}>
        <ResourceButton
          onClick={() => {
            history.push(href)
          }}
        >
          <StarredResourceLabel>{props.resource.name}</StarredResourceLabel>
        </ResourceButton>
        <StarButton onClick={onClick}>
          <StarIcon />
        </StarButton>
      </StarredResourceRoot>
    </TiltTooltip>
  )
}

export default function StarredResourceBar(props: StarredResourceBarProps) {
  return (
    <StarredResourceBarRoot>
      {props.resources.map((r) => (
        <StarredResource
          resource={r}
          key={r.name}
          unstar={props.unstar}
          isSelected={r.name === props.selectedResource}
        />
      ))}
    </StarredResourceBarRoot>
  )
}

// translates the view to a pared-down model so that `StarredResourceBar` can have a simple API for testing.
export function starredResourcePropsFromView(
  view: Proto.webviewView,
  selectedResource: string
): StarredResourceBarProps {
  const pinCtx = useSidebarPin()
  const namesAndStatuses = (view?.resources || []).flatMap((r) => {
    if (r.name && pinCtx.pinnedResources.includes(r.name)) {
      return [{ name: r.name, status: combinedStatus(r) }]
    } else {
      return []
    }
  })
  return {
    resources: namesAndStatuses,
    unstar: pinCtx.unpinResource,
    selectedResource: selectedResource,
  }
}
