package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCurrentDateTime(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_GetCurrentTimeUsec(m, l0, v5+int32(12), int32(0))
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_GetCurrentReplayRecPtr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+96)) = int32(1)
	if v7 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[186]))
		F_s_lock(m, v11+int32(96), int32(489255), int32(4609), int32(205720))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[186]))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(0)
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+48))
			if l0 != 0 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26
			} else {
			}
			return v25
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[186]))
		*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(0)
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+48))
		if l0 != 0 {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26
		} else {
		}
		return v25
	}
}
func F_GetCurrentRoleId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, _consts[949]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[946])))
	if v5 != 0 {
		v6 = v2
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_GetCurrentTimeUsec(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v7 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	v9 = *(*int64)(unsafe.Add(mBase, _consts[96]))
	v11 = *(*int64)(unsafe.Add(mBase, _consts[776]))
	if v9 == v11 {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[777]))
		if v7 == v14 {
			v32 = *(*int32)(unsafe.Add(mBase, _consts[778]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v32
			v35 = *(*int64)(unsafe.Add(mBase, _consts[779]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35
			v38 = *(*int64)(unsafe.Add(mBase, _consts[780]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
			v41 = *(*int64)(unsafe.Add(mBase, _consts[781]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
			v44 = *(*int64)(unsafe.Add(mBase, _consts[782]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v44
			v47 = *(*int64)(unsafe.Add(mBase, _consts[783]))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v47
			v50 = *(*int32)(unsafe.Add(mBase, _consts[784]))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
			if l2 != 0 {
				v53 = *(*int32)(unsafe.Add(mBase, _consts[785]))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
			} else {
			}
			return
		} else {
			v17 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[777])) = v17
			v23 = F_timestamp2tm(m, v9, int32(4470404), int32(4470356), int32(4470400), v17, v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if v23 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errmsg(m, int32(400033), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(495737), int32(432), int32(488533))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, _consts[776])) = v9
					v29 = *(*int32)(unsafe.Add(mBase, _consts[774]))
					*(*int32)(unsafe.Add(mBase, _consts[777])) = v29
					v32 = *(*int32)(unsafe.Add(mBase, _consts[778]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v32
					v35 = *(*int64)(unsafe.Add(mBase, _consts[779]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35
					v38 = *(*int64)(unsafe.Add(mBase, _consts[780]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
					v41 = *(*int64)(unsafe.Add(mBase, _consts[781]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
					v44 = *(*int64)(unsafe.Add(mBase, _consts[782]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v44
					v47 = *(*int64)(unsafe.Add(mBase, _consts[783]))
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v47
					v50 = *(*int32)(unsafe.Add(mBase, _consts[784]))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
					if l2 != 0 {
						v53 = *(*int32)(unsafe.Add(mBase, _consts[785]))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
					} else {
					}
					return
				}
			}
		}
	} else {
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[777])) = v17
		v23 = F_timestamp2tm(m, v9, int32(4470404), int32(4470356), int32(4470400), v17, v7)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_errmsg(m, int32(400033), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(495737), int32(432), int32(488533))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, _consts[776])) = v9
				v29 = *(*int32)(unsafe.Add(mBase, _consts[774]))
				*(*int32)(unsafe.Add(mBase, _consts[777])) = v29
				v32 = *(*int32)(unsafe.Add(mBase, _consts[778]))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v32
				v35 = *(*int64)(unsafe.Add(mBase, _consts[779]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v35
				v38 = *(*int64)(unsafe.Add(mBase, _consts[780]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
				v41 = *(*int64)(unsafe.Add(mBase, _consts[781]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
				v44 = *(*int64)(unsafe.Add(mBase, _consts[782]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v44
				v47 = *(*int64)(unsafe.Add(mBase, _consts[783]))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v47
				v50 = *(*int32)(unsafe.Add(mBase, _consts[784]))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v50
				if l2 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, _consts[785]))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
				} else {
				}
				return
			}
		}
	}
}
func F_GetCurrentVirtualXIDs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v17 = F_palloc(m, v14<<(uint(int32(3))%32))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		v26 = F_LWLockAcquire(m, v22+int32(512), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if int32(0) < v28 {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[603]))
				v40 = v3
				v41 = v3
				v43 = v34
				for {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v40<<(uint(int32(2))%32))))
					v52 = v43 + v49*int32(640)
					v54 = *(*int32)(unsafe.Add(mBase, _consts[91]))
					if v52 == v54 {
						v101 = v41
						v102 = v43
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, _consts[94]))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v40))))
						if v60&int32(7) != 0 {
							v101 = v41
							v102 = v43
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+60))
							v65 = *(*int32)(unsafe.Add(mBase, _consts[226]))
							if v63 != v65 {
								v101 = v41
								v102 = v43
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
								if v67 == int32(0) {
									v101 = v41
									v102 = v43
								} else {
									if l0 != 0 {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v67)) == int32(0) {
											v81 = base.B2i32(base.Ui32(v67) <= base.Ui32(l0))
										} else {
											v81 = base.B2i32(v67-l0 <= int32(0))
										}
										v83 = *(*int32)(unsafe.Add(mBase, _consts[603]))
										if v81 == int32(0) {
											v101 = v41
											v102 = v83
										} else {
											v87 = v83
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
											if v88 == int32(0) {
												v101 = v41
												v102 = v87
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
												v94 = v17 + v41<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v88
												*(*int32)(unsafe.Add(mBase, uint32(v94))) = v91
												v101 = v41 + int32(1)
												v102 = v87
											}
										}
									} else {
										v87 = v43
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
										if v88 == int32(0) {
											v101 = v41
											v102 = v87
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
											v94 = v17 + v41<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v88
											*(*int32)(unsafe.Add(mBase, uint32(v94))) = v91
											v101 = v41 + int32(1)
											v102 = v87
										}
									}
								}
							}
						}
					}
					v105 = v40 + int32(1)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v105 < v106 {
						v40 = v105
						v41 = v101
						v43 = v102
						continue
					} else {
						break
					}
					break
				}
				v114 = v101
			} else {
				v114 = v3
			}
			v120 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			F_LWLockRelease(m, v120+int32(512))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v114
				return v17
			}
		}
	}
}
func F_SetCurrentRoleId(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	if l0 == int32(0) {
		v6 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[946])) = uint8(v6)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[947]))
		if v9 == v6 {
			return
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[948])))
			v17 = v9
			v18 = v13
			*(*int32)(unsafe.Add(mBase, _consts[4])) = v17
			*(*int32)(unsafe.Add(mBase, _consts[949])) = v17
			if v18&int32(1) != 0 {
				v28 = int32(271632)
			} else {
				v28 = int32(336934)
			}
			F_SetConfigOption(m, int32(216545), v28, int32(0), int32(1))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v15 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[946])) = uint8(v15)
		v17 = l0
		v18 = l1
		*(*int32)(unsafe.Add(mBase, _consts[4])) = v17
		*(*int32)(unsafe.Add(mBase, _consts[949])) = v17
		if v18&int32(1) != 0 {
			v28 = int32(271632)
		} else {
			v28 = int32(336934)
		}
		F_SetConfigOption(m, int32(216545), v28, int32(0), int32(1))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			return
		}
	}
}
func F_SetCurrentStatementStartTimestamp(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v2 < int32(0) {
		v9 = m.G0
		v10 = int32(16)
		v11 = v9 - v10
		m.G0 = v11
		F___gettimeofday(m, v11)
		mBase = m.M
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+8)))
		m.G0 = v11 + v10
		*(*int64)(unsafe.Add(mBase, _consts[95])) = v15 + v14*int64(1000000) - int64(946684800000000)
	} else {
	}
	return
}
func F_current_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v4 != 0 {
		v5 = F_cstring_to_text(m, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
		return int32(0)
	}
}
func F_isCurrentGroup(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v88 int32
	_ = v88
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v19 = v17 - int32(1)
	if int32(0) <= v19 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L11
	} else {
		goto L26
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v121
L3:
	;
	v27 = v19
	goto L6
L4:
	;
	goto L5
L5:
	;
	v121 = int32(1)
	goto L2
L6:
	;
	v34 = v27 * int32(36)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34+v35)+32)))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v38 < v37 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	F_slot_getsomeattrs_int(m, l1, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v45 = v37 - int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v46))))
	v50 = v45 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50+v51)))
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v54 < v37 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L10
L13:
	;
	F_slot_getsomeattrs_int(m, l2, v37)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v58 = int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v45))))
	if v48&v58|v62&v58 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	if int32(0) < v27 {
		v27 = v27 - int32(1)
		goto L6
	} else {
		goto L25
	}
L18:
	;
	if v48 == v62&int32(255) {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v50)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v74 = v73 + v34
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v53
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+28)) = v72
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+16)) = uint8(v80)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = m.T0[v84].(func(*base.Module, int32) int32)(m, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L22
	}
L21:
	;
	v121 = int32(0)
	goto L2
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+16)))
	if v88 == int32(1) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v85 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v121 = int32(0)
	goto L2
L25:
	;
	goto L7
L26:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v130
	F_errmsg_internal(m, int32(529192), v14)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(489831), int32(258), int32(231377))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
