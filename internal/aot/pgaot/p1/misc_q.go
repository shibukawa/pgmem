package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_queue_listen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	goto L1
L1:
	;
	v14 = int32(4554128)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v18
	v20 = F_strlen(m, l1)
	mBase = m.M
	v23 = F_palloc(m, v20+int32(5))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l0
	v27 = v23 + int32(4)
	if (l1^v27)&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	if v103 != 0 {
		goto L27
	} else {
		goto L28
	}
L5:
	;
	goto L4
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v81)
	if v81&int32(255) == int32(0) {
		goto L5
	} else {
		goto L21
	}
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v80 = l1
	v81 = v33
	v82 = v27
	goto L6
L8:
	;
	goto L9
L9:
	;
	if l1&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = l1
	v39 = v27
	goto L13
L11:
	;
	v51 = l1
	v53 = v27
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v58 = int32(-2139062144)
	if (int32(16843008)-v55|v55)&v58 != v58 {
		v80 = v51
		v81 = v55
		v82 = v53
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v40)
	if v40 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v51 = v47
	v53 = v45
	goto L12
L15:
	;
	v44 = int32(1)
	v45 = v39 + v44
	v47 = v37 + v44
	if v47&int32(3) != 0 {
		v37 = v47
		v39 = v45
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v63 = v51
	v64 = v55
	v65 = v53
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v64
	v67 = int32(4)
	v68 = v65 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v71 = v63 + v67
	v75 = int32(-2139062144)
	if (v69|(int32(16843008)-v69))&v75 == v75 {
		v63 = v71
		v64 = v69
		v65 = v68
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v80 = v71
	v81 = v69
	v82 = v68
	goto L6
L20:
	;
	goto L19
L21:
	;
	v89 = v80
	v91 = v82
	goto L22
L22:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)) = uint8(v92)
	v94 = int32(1)
	if v92 != 0 {
		v89 = v89 + v94
		v91 = v91 + v94
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v15
	m.G0 = v9 + int32(16)
	return
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v126 = F_lappend(m, v125, v23)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L33
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v13 <= v104 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v109 = F_MemoryContextAlloc(m, v107, int32(12))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v23
	v117 = F_list_make1_impl(m, int32(1), v9+int32(8))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v117
	v120 = int32(4451084)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v121
	*(*int32)(unsafe.Add(mBase, _consts[151])) = v109
	goto L25
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v126
	goto L25
}
func F_quickdie(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v24 = int32(4461480)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[844]))
	*(*int32)(unsafe.Add(mBase, _consts[844])) = v25 | int32(4)
	F_sigprocmask(m, int32(4461480), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return
	} else {
		v35 = int32(4548780)
		v37 = *(*int32)(unsafe.Add(mBase, _consts[163]))
		v38 = int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[163])) = v37 + v38
		v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[147])))
		if v42 != v38 {
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, _consts[146]))
			if v46 != int32(2) {
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[146])) = int32(0)
			}
		}
		v53 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[337])) = v53
		v58 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
		if v58 != int32(1) {
			v66 = v53
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, _consts[571]))
			if v62 == int32(0) {
				v66 = v53
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+40))
				v66 = v65
			}
		}
		switch v66 {
		case 0:
			v70 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				if v70 == int32(0) {
					F__Exit(m, int32(2))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				} else {
					F_errcode(m, int32(16908741))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errmsg(m, int32(328195), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v118 = int32(3089)
							F_errfinish(m, int32(516919), v118, int32(417025))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								F__Exit(m, int32(2))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 1:
			v84 = F_errstart(m, int32(20), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				if v84 == int32(0) {
					F__Exit(m, int32(2))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				} else {
					F_errcode(m, int32(33685957))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errmsg(m, int32(137105), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							F_errdetail(m, int32(598312), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								F_errhint(m, int32(673232), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									v118 = int32(3101)
									F_errfinish(m, int32(516919), v118, int32(417025))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										F__Exit(m, int32(2))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		case 2:
			v106 = F_errstart(m, int32(20), int32(0))
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return
			} else {
				if v106 == int32(0) {
					F__Exit(m, int32(2))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				} else {
					F_errcode(m, int32(16908741))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						F_errmsg(m, int32(448338), int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							v118 = int32(3107)
							F_errfinish(m, int32(516919), v118, int32(417025))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								F__Exit(m, int32(2))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		default:
			F__Exit(m, int32(2))
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_quote_qualified_identifier(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	F_initStringInfo(m, v6+int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l0 != 0 {
			v14 = F_quote_identifier(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
				F_appendStringInfo(m, v6+int32(16), int32(633430), v6)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v24 = F_quote_identifier(m, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_appendStringInfoString(m, v6+int32(16), v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
							m.G0 = v6 + int32(32)
							return v28
						}
					}
				}
			}
		} else {
			v24 = F_quote_identifier(m, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_appendStringInfoString(m, v6+int32(16), v24)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
					m.G0 = v6 + int32(32)
					return v28
				}
			}
		}
	}
}
