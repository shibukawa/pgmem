package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bool_accum_inv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(335961), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(474012), int32(370), int32(30312))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(335961), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474012), int32(370), int32(30312))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v7 != 0 {
			} else {
				v8 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
				*(*int64)(unsafe.Add(mBase, uint32(v4))) = v8 - int64(1)
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v12 == int32(0) {
				} else {
					v15 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v15 - int64(1)
				}
			}
			return v4
		}
	}
}
func F_bool_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(l1 != v4))
	return base.B2i32(l1 == v4)
}
func F_call_bool_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v14 == int32(0) {
		v81 = v13
		m.G0 = v11 - int32(-64)
		return v81
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[249])) = int32(50856066)
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[422])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[251])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[252])) = v21
		v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v81 = v13
				m.G0 = v11 - int32(-64)
				return v81
			} else {
				v34 = F_errstart(m, l4, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != 0 {
						v37 = *(*int32)(unsafe.Add(mBase, _consts[249]))
						F_errcode(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _consts[422]))
							if v41 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v41
								F_errmsg_internal(m, int32(195784), v9+int32(-16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[251]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(195784), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(195784), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(476372), int32(6836), int32(299702))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(476372), int32(6836), int32(299702))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(195784), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(476372), int32(6836), int32(299702))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(476372), int32(6836), int32(299702))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v48
								F_errmsg(m, int32(465420), v9+int32(-32))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[251]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(195784), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(195784), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(476372), int32(6836), int32(299702))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(476372), int32(6836), int32(299702))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[252]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(195784), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(476372), int32(6836), int32(299702))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(476372), int32(6836), int32(299702))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_FlushErrorState(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v81 = int32(0)
							m.G0 = v11 - int32(-64)
							return v81
						}
					}
				}
			}
		}
	}
}
func F_parse_bool(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v61 - int32(48) {
	case 0:
		goto L25
	case 1:
		goto L26
	default:
		goto L23
	case 22, 54:
		goto L30
	case 30, 62:
		goto L28
	case 31, 63:
		goto L27
	case 36, 68:
		goto L31
	case 41, 73:
		goto L29
	}
L2:
	;
	v59 = v51 - l0
	goto L1
L3:
	;
	v30 = v26
	goto L12
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v15 = l0
	goto L8
L8:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v51 = v19
	goto L2
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v45 = v30
	goto L15
L14:
	;
	goto L13
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v51 = v45
	goto L2
L17:
	;
	goto L16
L18:
	;
	return v122
L19:
	;
	v122 = v116
	goto L18
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v112)
	v116 = v114
	goto L19
L21:
	;
	if l1 == int32(0) {
		v116 = v107
		goto L19
	} else {
		goto L48
	}
L22:
	;
	v107 = int32(1)
	goto L21
L23:
	;
	v103 = int32(0)
	if l1 != 0 {
		v112 = v103
		v114 = v103
		goto L20
	} else {
		goto L47
	}
L24:
	;
	v112 = int32(0)
	v114 = v98
	goto L20
L25:
	;
	v93 = int32(1)
	if v59 != v93 {
		goto L23
	} else {
		goto L45
	}
L26:
	;
	v90 = int32(1)
	if v59 != v90 {
		goto L23
	} else {
		goto L44
	}
L27:
	;
	v79 = int32(2)
	if base.Ui32(v59) <= base.Ui32(v79) {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v75 = F_pg_strncasecmp(m, l0, int32(228548), v59)
	mBase = m.M
	if v75 != 0 {
		goto L23
	} else {
		goto L36
	}
L29:
	;
	v71 = F_pg_strncasecmp(m, l0, int32(147524), v59)
	mBase = m.M
	if v71 == int32(0) {
		goto L22
	} else {
		goto L35
	}
L30:
	;
	v67 = F_pg_strncasecmp(m, l0, int32(343932), v59)
	mBase = m.M
	if v67 != 0 {
		goto L23
	} else {
		goto L33
	}
L31:
	;
	v65 = F_pg_strncasecmp(m, l0, int32(327357), v59)
	mBase = m.M
	if v65 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L22
L33:
	;
	if l1 != 0 {
		v98 = int32(1)
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v122 = int32(1)
	goto L18
L35:
	;
	goto L23
L36:
	;
	if l1 != 0 {
		v98 = int32(1)
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v122 = int32(1)
	goto L18
L38:
	;
	v82 = v79
	goto L40
L39:
	;
	v82 = v59
	goto L40
L40:
	;
	v83 = F_pg_strncasecmp(m, l0, int32(259589), v82)
	mBase = m.M
	if v83 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v87 = F_pg_strncasecmp(m, l0, int32(322164), v82)
	mBase = m.M
	if v87 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	if l1 != 0 {
		v98 = int32(1)
		goto L24
	} else {
		goto L43
	}
L43:
	;
	v122 = int32(1)
	goto L18
L44:
	;
	v107 = v90
	goto L21
L45:
	;
	if l1 != 0 {
		v98 = v93
		goto L24
	} else {
		goto L46
	}
L46:
	;
	v122 = int32(1)
	goto L18
L47:
	;
	v116 = v103
	goto L19
L48:
	;
	v112 = v107
	v114 = int32(1)
	goto L20
}
