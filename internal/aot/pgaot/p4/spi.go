package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_cursor_close(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	if l0 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(270299), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errfinish(m, int32(518140), int32(1865), int32(376205))
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		F_PortalDrop(m, l0, int32(0))
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
}
func F_SPI_execute(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v2 = l1
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = int32(-6)
	if l0 == int32(0) {
		v77 = v11
		m.G0 = v9 - int32(-64)
		return v77
	} else {
		if l2 < int32(0) {
			v77 = v11
			m.G0 = v9 - int32(-64)
			return v77
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[217]))
			if v17 == int32(0) {
				v77 = int32(-4)
				m.G0 = v9 - int32(-64)
				return v77
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[25]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				v25 = *(*int32)(unsafe.Add(mBase, _consts[217]))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v23
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
				v29 = v7 + int32(-20)
				v30 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v27
				*(*int64)(unsafe.Add(mBase, uint32(v9)+36)) = v30
				*(*int64)(unsafe.Add(mBase, uint32(v9)+52)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(2048)
				*(*int64)(unsafe.Add(mBase, uint32(v9)+28)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(569278163)
				F__SPI_prepare_oneshot_plan(m, l0, v7+int32(-40))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v52
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = base.I64_extend_i32_u(l2)
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v52
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)) = uint8(v2)
					v61 = int32(0)
					v64 = F__SPI_execute_plan(m, v7+int32(-40), v9, v61, v61, int32(1))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, _consts[217]))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = int32(0)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
						F_MemoryContextReset(m, v73)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v77 = v64
							m.G0 = v9 - int32(-64)
							return v77
						}
					}
				}
			}
		}
	}
}
func F_SPI_execute_plan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = int32(-6)
	if l0 == v4 {
		v130 = v17
		m.G0 = v15 + int32(32)
		return v130
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v20 != int32(569278163) {
			v130 = v17
			m.G0 = v15 + int32(32)
			return v130
		} else {
			if l1 != 0 {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[217]))
				if v28 == int32(0) {
					v130 = int32(-4)
					m.G0 = v15 + int32(32)
					return v130
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[25]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
					v36 = *(*int32)(unsafe.Add(mBase, _consts[217]))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v34
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v39
					v41 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v41
					*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v41
					*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v41
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if int32(0) < v47 {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v51 = F_makeParamList(m, v47)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v61 = int32(0)
							for {
								v72 = v51 + int32(32) + v61*int32(12)
								v74 = v61 << (uint(int32(2)) % 32)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1+v74)))
								*(*int32)(unsafe.Add(mBase, uint32(v72))) = v76
								if l2 != 0 {
									v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v61))))
									v83 = base.B2i32(v80 == int32(110))
								} else {
									v83 = int32(0)
								}
								v84 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v84)
								*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)) = uint8(v83)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v50+v74)))
								*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v88
								v91 = v61 + v84
								if v91 != v47 {
									v61 = v91
									continue
								} else {
									break
								}
								break
							}
							v102 = v51
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v105)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v102
							*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(0)
							v112 = int32(0)
							v115 = F__SPI_execute_plan(m, l0, v15+int32(8), v112, v112, v105)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v119 = *(*int32)(unsafe.Add(mBase, _consts[217]))
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = int32(0)
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
								F_MemoryContextReset(m, v124)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v130 = v115
									m.G0 = v15 + int32(32)
									return v130
								}
							}
						}
					} else {
						v102 = v4
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v105)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v102
						*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(0)
						v112 = int32(0)
						v115 = F__SPI_execute_plan(m, l0, v15+int32(8), v112, v112, v105)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v119 = *(*int32)(unsafe.Add(mBase, _consts[217]))
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = int32(0)
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
							F_MemoryContextReset(m, v124)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								v130 = v115
								m.G0 = v15 + int32(32)
								return v130
							}
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v23 <= int32(0) {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[217]))
					if v28 == int32(0) {
						v130 = int32(-4)
						m.G0 = v15 + int32(32)
						return v130
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _consts[25]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
						v36 = *(*int32)(unsafe.Add(mBase, _consts[217]))
						*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v34
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v39
						v41 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v41
						*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v41
						*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v41
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if int32(0) < v47 {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v51 = F_makeParamList(m, v47)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v61 = int32(0)
								for {
									v72 = v51 + int32(32) + v61*int32(12)
									v74 = v61 << (uint(int32(2)) % 32)
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l1+v74)))
									*(*int32)(unsafe.Add(mBase, uint32(v72))) = v76
									if l2 != 0 {
										v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v61))))
										v83 = base.B2i32(v80 == int32(110))
									} else {
										v83 = int32(0)
									}
									v84 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v84)
									*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)) = uint8(v83)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v50+v74)))
									*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v88
									v91 = v61 + v84
									if v91 != v47 {
										v61 = v91
										continue
									} else {
										break
									}
									break
								}
								v102 = v51
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v105)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v102
								*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(0)
								v112 = int32(0)
								v115 = F__SPI_execute_plan(m, l0, v15+int32(8), v112, v112, v105)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, _consts[217]))
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
									*(*int32)(unsafe.Add(mBase, _consts[28])) = v120
									*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = int32(0)
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
									F_MemoryContextReset(m, v124)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v130 = v115
										m.G0 = v15 + int32(32)
										return v130
									}
								}
							}
						} else {
							v102 = v4
							v105 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v105)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v102
							*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = int64(0)
							v112 = int32(0)
							v115 = F__SPI_execute_plan(m, l0, v15+int32(8), v112, v112, v105)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v119 = *(*int32)(unsafe.Add(mBase, _consts[217]))
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = int32(0)
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
								F_MemoryContextReset(m, v124)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v130 = v115
									m.G0 = v15 + int32(32)
									return v130
								}
							}
						}
					}
				} else {
					v130 = int32(-7)
					m.G0 = v15 + int32(32)
					return v130
				}
			}
		}
	}
}
func F_SPI_fnumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v13 + int32(1)
L2:
	;
	v13 = v3
	v14 = v6
	goto L5
L3:
	;
	goto L4
L4:
	;
	v53 = F_strcmp(m, int32(793692), l1)
	mBase = m.M
	if v53 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v16 = int32(4)
	v21 = l0 + int32(20) + v14<<(uint(v16)%32) + v13*int32(100)
	v23 = v21 + v16
	if v23|l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	if v37 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v29 = int32(-1)
	goto L10
L9:
	;
	v29 = int32(0)
	goto L10
L10:
	;
	if v23 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = int32(1)
	goto L13
L12:
	;
	v30 = v29
	goto L13
L13:
	;
	if v23 == int32(0) {
		v37 = v30
		goto L14
	} else {
		goto L15
	}
L14:
	;
	goto L7
L15:
	;
	if l1 == int32(0) {
		v37 = v30
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v36 = F_strncmp(m, v23, l1, int32(64))
	mBase = m.M
	v37 = v36
	goto L14
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+91)))
	if v40 != int32(1) {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v44 = v13 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v44 < v45 {
		v13 = v44
		v14 = v45
		goto L5
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	goto L6
L22:
	;
	if v82 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	v82 = int32(793688)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v58 = F_strcmp(m, int32(793792), l1)
	mBase = m.M
	if v58 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v82 = int32(793788)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v63 = F_strcmp(m, int32(793892), l1)
	mBase = m.M
	if v63 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v82 = int32(793888)
	goto L22
L30:
	;
	goto L31
L31:
	;
	v68 = F_strcmp(m, int32(793992), l1)
	mBase = m.M
	if v68 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v82 = int32(793988)
	goto L22
L33:
	;
	goto L34
L34:
	;
	v73 = F_strcmp(m, int32(794092), l1)
	mBase = m.M
	if v73 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v82 = int32(794088)
	goto L22
L36:
	;
	goto L37
L37:
	;
	v80 = F_strcmp(m, int32(794192), l1)
	mBase = m.M
	if v80 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v81 = int32(0)
	goto L40
L39:
	;
	v81 = int32(794188)
	goto L40
L40:
	;
	v82 = v81
	goto L22
L41:
	;
	return int32(-9)
L42:
	;
	goto L43
L43:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+74)))
	return v87
}
func F_SPI_freetuptable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var __phi28 int32
	_ = __phi28
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v68 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L20
	} else {
		goto L22
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v17 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if l0 == v17-int32(28) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v42
	v46 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v47 == l0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v42 = v20
	v43 = v14 + int32(16)
	goto L6
L8:
	;
	goto L9
L9:
	;
	__phi28 = v17
	__phi30 = v20
	v28 = __phi28
	v30 = __phi30
	goto L10
L10:
	;
	if v30 == int32(0) {
		goto L3
	} else {
		goto L12
	}
L11:
	;
	v42 = v34
	v43 = v28
	goto L6
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v30-int32(28) != l0 {
		__phi28 = v30
		__phi30 = v34
		v28 = __phi28
		v30 = __phi30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = int32(0)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	if v52 == l0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[511])) = int32(0)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_MemoryContextDelete(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	goto L1
L22:
	;
	if v68 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(249345), v9)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(518140), int32(1424), int32(407489))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L1
}
func F_SPI_getvalue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, _consts[512])) = v4
	if l2 < int32(-6) {
		*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-9)
		v55 = v4
		m.G0 = v8 + int32(16)
		return v55
	} else {
		if l2 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-9)
			v55 = v4
			m.G0 = v8 + int32(16)
			return v55
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if l2 <= v17 {
				v24 = F_heap_getattr_2(m, l0, l2, l1, v8+int32(15))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v28 != 0 {
						v55 = v4
						m.G0 = v8 + int32(16)
						return v55
					} else {
						if int32(0) < l2 {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v43 = l1 + v31<<(uint(int32(4))%32) + l2*int32(100) - int32(80)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
							F_getTypeOutputInfo(m, v44, v8+int32(8), v8+int32(7))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
								v52 = F_OidOutputFunctionCall(m, v51, v24)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v55 = v52
									m.G0 = v8 + int32(16)
									return v55
								}
							}
						} else {
							v41 = F_SystemAttributeDefinition(m, base.I32_extend16_s(l2))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = v41
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
								F_getTypeOutputInfo(m, v44, v8+int32(8), v8+int32(7))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
									v52 = F_OidOutputFunctionCall(m, v51, v24)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v55 = v52
										m.G0 = v8 + int32(16)
										return v55
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-9)
				v55 = v4
				m.G0 = v8 + int32(16)
				return v55
			}
		}
	}
}
func F_SPI_plan_get_cached_plan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v9 != 0 {
		v51 = v2
		m.G0 = v7 + int32(32)
		return v51
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v10 == int32(0) {
			v51 = v2
			m.G0 = v7 + int32(32)
			return v51
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if v13 != int32(1) {
				v51 = v2
				m.G0 = v7 + int32(32)
				return v51
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(779)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v20
				v24 = int32(4541928)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[394]))
				*(*int32)(unsafe.Add(mBase, _consts[394])) = v7 + int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v7 + int32(24)
				v34 = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[182]))
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
				if v38 != 0 {
					v39 = v36
				} else {
					v39 = v34
				}
				v41 = *(*int32)(unsafe.Add(mBase, _consts[217]))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
				v43 = F_GetCachedPlan(m, v17, v34, v39, v42)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, _consts[394])) = v48
					v51 = v43
					m.G0 = v7 + int32(32)
					return v51
				}
			}
		}
	}
}
func F_SPI_prepare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-6)
		v72 = int32(0)
		m.G0 = v9 + int32(48)
		return v72
	} else {
		if l1 < int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-6)
			v72 = int32(0)
			m.G0 = v9 + int32(48)
			return v72
		} else {
			if l1 == int32(0) {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[217]))
				if v22 != 0 {
					v24 = *(*int32)(unsafe.Add(mBase, _consts[25]))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					v27 = *(*int32)(unsafe.Add(mBase, _consts[217]))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v25
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v30
					v32 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = v32
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[512])) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = v32
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(569278163)
					F__SPI_prepare_plan(m, l0, v9+int32(8))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v55 = F__SPI_make_plan_non_temp(m, v9+int32(8))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, _consts[217]))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v60
							*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(0)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
							F_MemoryContextReset(m, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v72 = v55
								m.G0 = v9 + int32(48)
								return v72
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-4)
					v72 = int32(0)
					m.G0 = v9 + int32(48)
					return v72
				}
			} else {
				if l2 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[217]))
					if v22 != 0 {
						v24 = *(*int32)(unsafe.Add(mBase, _consts[25]))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						v27 = *(*int32)(unsafe.Add(mBase, _consts[217]))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v25
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v30
						v32 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = v32
						v35 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[512])) = v35
						*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = v32
						*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(569278163)
						F__SPI_prepare_plan(m, l0, v9+int32(8))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v55 = F__SPI_make_plan_non_temp(m, v9+int32(8))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, _consts[217]))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v60
								*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(0)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
								F_MemoryContextReset(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v72 = v55
									m.G0 = v9 + int32(48)
									return v72
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-4)
						v72 = int32(0)
						m.G0 = v9 + int32(48)
						return v72
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[512])) = int32(-6)
					v72 = int32(0)
					m.G0 = v9 + int32(48)
					return v72
				}
			}
		}
	}
}
func F__SPI_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 == int32(0) {
		m.G0 = v7 + int32(48)
		return
	} else {
		v12 = F_geterrposition(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if int32(0) < v12 {
				v17 = F_errposition(m, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_internalerrposition(m, v12)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v21 = F_internalerrquery(m, v9)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				switch v23 - int32(2) {
				case 0:
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v9
						F_errcontext_msg(m, int32(737534), v7+int32(16))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					}
				case 1, 2, 3:
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v9
						F_errcontext_msg(m, int32(727260), v7+int32(32))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					}
				default:
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
						F_errcontext_msg(m, int32(727375), v7)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							m.G0 = v7 + int32(48)
							return
						}
					}
				}
			}
		}
	}
}
func F__SPI_prepare_oneshot_plan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int64
	_ = v64
	var v73 int32
	_ = v73
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l0
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(779)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v14
	v18 = int32(4541928)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v11 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v11 + int32(24)
	v28 = F_raw_parser(m, l0, v14)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v110
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v117
	m.G0 = v11 + int32(32)
	return
L2:
	;
	return
L3:
	;
	if v28 == int32(0) {
		v110 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v33 <= v32 {
		v110 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v40 = v32
	v41 = v3
	goto L6
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v50 = F_CreateCommandTag(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v110 = v99
	goto L1
L8:
	;
	v53 = F_palloc0(m, int32(144))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+52)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(195726186)
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+20)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v53)+28)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v53)+36)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v53)+41)) = v64
	v73 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v53)+56)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v53)+60)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v53)+68)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v53)+76)) = v64
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+84)) = uint16(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+93)) = v55
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+92)) = uint8(v87)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+96)) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v53)+120)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v53)+112)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+128)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v53)+136)) = v64
	v99 = F_lappend(m, v41, v53)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v102 = v40 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v102 < v103 {
		v40 = v102
		v41 = v99
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
}
