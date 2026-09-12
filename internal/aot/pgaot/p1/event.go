package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddWaitEventToSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	if l1 == int32(32) {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v10)
		v13 = int32(16)
	} else {
		v13 = l1
	}
	if l3 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		v16 = *(*int32)(unsafe.Add(mBase, _consts[156]))
		if v14 != v16 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(127308), int32(0))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return
				} else {
					F_errfinish(m, int32(485692), int32(587), int32(106851))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v18 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(320808), int32(0))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						F_errfinish(m, int32(485692), int32(589), int32(106851))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if v13&int32(1) != 0 {
					if v13&int32(134) != 0 {
						v41 = base.B2i32(l2 == int32(-1))
					} else {
						v41 = int32(0)
					}
					if v41 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(105879), int32(0))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								F_errfinish(m, int32(485692), int32(601), int32(106851))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v42 = int32(1)
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 + v42
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v50 = v47 + v43<<(uint(int32(4))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v13
						switch v13 - v42 {
						case 0:
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v60
							v63 = int32(4089264)
							v64 = v60
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v65
							v68 = v64
						default:
							v68 = v43
						case 15:
							v63 = int32(4088788)
							v64 = v43
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v65
							v68 = v64
						}
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v72 = v69 + v68<<(uint(int32(3))%32)
						v73 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v73)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v72))) = v75
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						switch v77 - int32(1) {
						case 0, 15:
							v105 = v42
							*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v105)
						default:
							v80 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v80)
							v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
							v83 = int32(1)
							v86 = int32(base.Ui32(v82)>>(uint(v83)%32)) & v83
							*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v86)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							if v88&int32(4) != 0 {
								v92 = v86 | int32(4)
								*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v92)
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								v95 = v92
								v96 = v94
							} else {
								v95 = v86
								v96 = v88
							}
							if v96&int32(128) == int32(0) {
							} else {
								v105 = v95 | int32(8192)
								*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v105)
							}
						}
						return
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(104877), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_errfinish(m, int32(485692), int32(591), int32(106851))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		if v13&int32(1) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(320843), int32(0))
				mBase = m.M
				v144 = m.ExcPending
				if v144 != 0 {
					return
				} else {
					F_errfinish(m, int32(485692), int32(596), int32(106851))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if v13&int32(134) != 0 {
				v41 = base.B2i32(l2 == int32(-1))
			} else {
				v41 = int32(0)
			}
			if v41 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(105879), int32(0))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						F_errfinish(m, int32(485692), int32(601), int32(106851))
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v42 = int32(1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 + v42
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v50 = v47 + v43<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v50))) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v13
				switch v13 - v42 {
				case 0:
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v60
					v63 = int32(4089264)
					v64 = v60
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v65
					v68 = v64
				default:
					v68 = v43
				case 15:
					v63 = int32(4088788)
					v64 = v43
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v65
					v68 = v64
				}
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v72 = v69 + v68<<(uint(int32(3))%32)
				v73 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v73)
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v72))) = v75
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
				switch v77 - int32(1) {
				case 0, 15:
					v105 = v42
					*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v105)
				default:
					v80 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v80)
					v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
					v83 = int32(1)
					v86 = int32(base.Ui32(v82)>>(uint(v83)%32)) & v83
					*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v86)
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
					if v88&int32(4) != 0 {
						v92 = v86 | int32(4)
						*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v92)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						v95 = v92
						v96 = v94
					} else {
						v95 = v86
						v96 = v88
					}
					if v96&int32(128) == int32(0) {
					} else {
						v105 = v95 | int32(8192)
						*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v105)
					}
				}
				return
			}
		}
	}
}
func F_CreateWaitEventSet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v6 = l1 << (uint(int32(4)) % 32)
	if l0 != 0 {
		F_ResourceOwnerEnlarge(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[182]))
			v18 = F_MemoryContextAllocZero(m, v17, v6+l1<<(uint(int32(3))%32)+int32(32))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v20
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)) = uint8(v20)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
				v26 = v18 + int32(32)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v26 + v6
				if l0 != 0 {
					F_ResourceOwnerRemember(m, l0, v18, int32(1604396))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
						return v18
					}
				} else {
					return v18
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[182]))
		v18 = F_MemoryContextAllocZero(m, v17, v6+l1<<(uint(int32(3))%32)+int32(32))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v20
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
			v26 = v18 + int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v26 + v6
			if l0 != 0 {
				F_ResourceOwnerRemember(m, l0, v18, int32(1604396))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
					return v18
				}
			} else {
				return v18
			}
		}
	}
}
func F_EventTriggerAlterTableEnd(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	if v6 == int32(0) {
		return
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if v12 != 0 {
				v13 = int32(4470752)
				v14 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v16
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
				v19 = F_lappend(m, v18, v10)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[364]))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v19
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v14
					v30 = v22
					*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v11
					return
				}
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[364]))
					v30 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v11
					return
				}
			}
		}
	}
}
func F_EventTriggerCollectSimpleCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v7 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	if v7 == int32(0) {
		return
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
		if v10 != 0 {
			return
		} else {
			v11 = int32(4470752)
			v12 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v14
			v17 = F_palloc(m, int32(40))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
				v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[295])))
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v22)
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+12)) = v24
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v26
				v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v30
				v32 = F_copyObjectImpl(m, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v32
					v36 = *(*int32)(unsafe.Add(mBase, _consts[364]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
					v38 = F_lappend(m, v37, v17)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _consts[364]))
						*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v38
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
						return
					}
				}
			}
		}
	}
}
func F_EventTriggerSQLDrop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v18 = v2
	v19 = v2
	v20 = v2
	v21 = v2
	v22 = v2
	v23 = v2
	v24 = v13
	v25 = int32(-1)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v25 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v144 = int32(m.ExcTag)
	v145 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v144 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v85
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v84
	v127 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v83
	F_pg_re_throw(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		v143 = v87
		goto L5
	} else {
		goto L25
	}
L7:
	;
	m.G0 = v13 + int32(32)
	return
L8:
	;
	v29 = v24 - int32(16)
	m.G0 = v29
	v32 = v29 - int32(160)
	m.G0 = v32
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	if v35 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v81 = v18
	v82 = v19
	v83 = v20
	v84 = v21
	v85 = v22
	v86 = v23
	v87 = v24
	goto L10
L10:
	;
	if v86 != 0 {
		goto L6
	} else {
		goto L22
	}
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[365])))
	if v39 != int32(1) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	if v43 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v46 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	v56 = F_EventTriggerCommonSetup(m, l0, int32(2), int32(231112), v29)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		v143 = v32
		goto L5
	} else {
		goto L15
	}
L15:
	;
	if v56 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_CommandCounterIncrement(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		v143 = v32
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+8)) = uint8(v70)
	v73 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v75 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v13 + int32(8)
	goto L21
L19:
	;
	v81 = v56
	v82 = v32
	v83 = v29
	v84 = v75
	v85 = v73
	v86 = int32(0)
	v87 = v32
	goto L10
L21:
	;
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v83
	F_EventTriggerInvoke(m, v81, v83)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		v143 = v87
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+8)) = uint8(v99)
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v85
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v83
	F_list_free(m, v81)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		v143 = v87
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L7
L25:
	;
	goto L4
L26:
	;
	v149 = int32(v145)
	m.G0 = v143
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	if v13+int32(8) == v156 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	m.ExcPending = 1
	goto L35
L28:
	;
	if v159 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v159 = v158
	goto L31
L30:
	;
	v159 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v18 = v162
	v19 = v161
	v20 = v160
	v21 = v163
	v22 = v164
	v23 = v151
	v24 = v143
	v25 = v159
	goto L1
L33:
	;
	goto L34
L34:
	;
	F___wasm_longjmp(m, v152, v151)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_EventTriggerSupportsObjectType(m *base.Module, l0 int32) int32 {
	return (base.I32_wrap_i64(int64(base.Ui64(int64(4389322341887))>>(uint(base.I64_extend_i32_u(l0))%64))) | base.B2i32(base.Ui32(int32(42)) < base.Ui32(l0))) & int32(1)
}
func F_FreeWaitEventSet(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		F_ResourceOwnerForget(m, v3, l0, int32(1604396))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
