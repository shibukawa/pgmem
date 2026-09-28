package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckRelationTableSpaceMove(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if l1 == v12 {
		v34 = v3
		m.G0 = v9 + int32(16)
		return v34
	} else {
		if v12 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRelationTableSpaceMove[0]))
			if l1 == v17 {
				v34 = v3
				m.G0 = v9 + int32(16)
				return v34
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
				switch v19 - int32(83) {
				case 0, 22, 26, 31, 33:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
					if v22 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48 + int32(4)
								F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_0), v9)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3747), int32(_a_F_CheckRelationTableSpaceMove_2))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if l1 == int32(1664) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3753), int32(_a_F_CheckRelationTableSpaceMove_2))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v27 = int32(1)
							v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
							if v28 != int32(116) {
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							} else {
								v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v31 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3762), int32(_a_F_CheckRelationTableSpaceMove_2))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v34 = v27
									m.G0 = v9 + int32(16)
									return v34
								}
							}
						}
					}
				default:
					if l1 == int32(1664) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3753), int32(_a_F_CheckRelationTableSpaceMove_2))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v27 = int32(1)
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
						if v28 != int32(116) {
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v31 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3762), int32(_a_F_CheckRelationTableSpaceMove_2))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
			switch v19 - int32(83) {
			case 0, 22, 26, 31, 33:
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48 + int32(4)
							F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_0), v9)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3747), int32(_a_F_CheckRelationTableSpaceMove_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if l1 == int32(1664) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3753), int32(_a_F_CheckRelationTableSpaceMove_2))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v27 = int32(1)
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
						if v28 != int32(116) {
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v31 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3762), int32(_a_F_CheckRelationTableSpaceMove_2))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			default:
				if l1 == int32(1664) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3753), int32(_a_F_CheckRelationTableSpaceMove_2))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v27 = int32(1)
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
					if v28 != int32(116) {
						v34 = v27
						m.G0 = v9 + int32(16)
						return v34
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v31 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3762), int32(_a_F_CheckRelationTableSpaceMove_2))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(224)
	m.G0 = v10
	switch l2 - int32(1663) {
	case 0:
		if l4 == int32(-1) {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = l1
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v40
				v45 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_0), v10+int32(176))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = l1
				v52 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_1), v10+int32(160))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		} else {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+216)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+212)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = l1
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+220)) = v61
				v66 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_2), v10+int32(208))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = l1
				v74 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_3), v10+int32(192))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		}
	case 1:
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = l3
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+148)) = v19
			v24 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_4), v10+int32(144))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v10 + int32(224)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = l3
			v30 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_5), v10+int32(128))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v10 + int32(224)
				return
			}
		}
	default:
		if l4 == int32(-1) {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l3
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a_F_GetRelationPath_7)
				v94 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_8), v10+int32(32))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_GetRelationPath_7)
				v104 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_9), v10)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		} else {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = l4
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v112
				*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = int32(_a_F_GetRelationPath_7)
				v123 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_10), v10+int32(96))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(_a_F_GetRelationPath_7)
				v136 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_11), v10-int32(-64))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		}
	}
}
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v53 int32
	_ = v53
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v82 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v82 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v82 = int32(0)
				} else {
					v82 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v23))|base.B2i32(int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v82 = v11
				} else {
					v82 = int32(0)
				}
			} else {
				v82 = v11
			}
		}
	} else {
		if l0 <= int32(_a_F_LockRelationOid_0) {
			v36 = l0 - int32(_a_F_LockRelationOid_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v36))|base.B2i32(int32(1)<<(uint(v36)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v82 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v82 = int32(0)
					} else {
						v82 = v11
					}
				}
			} else {
				v82 = v11
			}
		} else {
			switch l0 - int32(_a_F_LockRelationOid_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v82 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v82 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_LockRelationOid_3)) < base.Ui32(int32(3)) {
					v82 = v11
				} else {
					v53 = l0 - int32(_a_F_LockRelationOid_4)
					if base.Ui32(int32(15)) < base.Ui32(v53) {
						v82 = int32(0)
					} else {
						if int32(1)<<(uint(v53)%32)&int32(_a_F_LockRelationOid_5) != 0 {
							v82 = v11
						} else {
							v82 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelationOid[0]))
	if v82 != 0 {
		v89 = int32(0)
	} else {
		v89 = v88
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v89
	v93 = int32(0)
	v98 = F_LockAcquireExtended(m, v7+int32(16), l1, v93, v93, v7+int32(12), v93)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		return
	} else {
		if v98 != int32(3) {
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v105 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v104)+53)) = uint8(v105)
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	}
}
func F_RelationBuildDesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v78 int32
	_ = v78
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int64
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int64
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int64
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int64
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int64
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v902 int32
	_ = v902
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1039 int64
	_ = v1039
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	v21 = m.G0
	v23 = v21 - int32(352)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0]))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[2]))
	if v30 <= v28 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = F_repalloc(m, v26, v30<<(uint(int32(4))%32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v46 = v28
	v47 = v26
	goto L3
L3:
	;
	v49 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1])) = v46 + v49
	v53 = v46 << (uint(int32(3)) % 32)
	v54 = v47 + v53
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)) = uint8(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = l0
	v60 = F_ScanPgRelation(m, l0, v49, v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[2])) = v30 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0])) = v34
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	v46 = v45
	v47 = v34
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L4
	} else {
		goto L275
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L4
	} else {
		goto L272
	}
L8:
	;
	m.G0 = v23 + int32(352)
	return v1202
L9:
	;
	v1145 = int32(_a_F_RelationBuildDesc_0)
	v1147 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1])) = v1147 - int32(1)
	if l1 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L10:
	;
	if v60 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v78 = v60
	goto L14
L12:
	;
	goto L13
L13:
	;
	v1139 = int32(_a_F_RelationBuildDesc_0)
	v1141 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1])) = v1141 - int32(1)
	v1202 = int32(0)
	goto L8
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+22)))
	v86 = v84 + v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v88 = int32(_a_F_RelationBuildDesc_1)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3]))
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v92
	v95 = F_palloc0(m, int32(276))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
	v100 = F_palloc(m, int32(144))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	base.MemoryCopy(m, v100, v86, int32(144))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = v100
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+120)))
	v106 = F_CreateTemplateTupleDesc(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+52)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v106)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v89
	v113 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v95)+32)) = v113
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+25)) = uint8(v115)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v95)+56)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = v113
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+118)))
	switch v123 - int32(112) {
	case 0, 5:
		goto L20
	default:
		goto L21
	case 4:
		goto L22
	}
L19:
	;
	v178 = v95 + int32(56)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+72))
	if v181 != 0 {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v171)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = int32(-1)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L37
	}
L22:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+68))
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[5]))
	if v130 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v138 != 0 {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	v131 = int32(1)
	if v126 == v130 {
		v138 = v131
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v138 = int32(0)
	goto L24
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[6]))
	if v134 == v126 {
		v138 = v131
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[7]))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[8]))
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v143)
	if v142 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v151 = F_GetTempNamespaceProcNumber(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L36
	}
L33:
	;
	v147 = v140
	goto L35
L34:
	;
	v147 = v142
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = v147
	goto L19
L36:
	;
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v153)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = v151
	goto L19
L37:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v161 = int32(*(*int8)(unsafe.Add(mBase, uint32(v160)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v161
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_2), v23)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(1197), int32(_a_F_RelationBuildDesc_4))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v183 = v181
	goto L42
L41:
	;
	v183 = int32(2249)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+8)) = int32(-1)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v192 = F_MemoryContextAllocZero(m, v190, int32(20))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v195 = v23 + int32(160)
	v199 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v95)+56)))
	F_ScanKeyInit(m, v195, int32(1), int32(3), int32(184), v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v202 = int32(5)
	F_ScanKeyInit(m, v23+int32(216), v202, v202, int32(146), int64(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v210 = F_table_open(m, int32(1249), int32(1))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationBuildDesc[9])))
	v217 = F_systable_beginscan(m, v210, int32(2659), v214, int32(0), int32(2), v195)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v220 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219)+120)))
	v227 = v220
	v228 = int32(0)
	v232 = int32(0)
	goto L49
L48:
	;
	F_systable_endscan(m, v217)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L85
	}
L49:
	;
	v242 = F_systable_getnext(m, v217)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	v377 = int32(0)
	v378 = v372
	v381 = v369
	goto L48
L51:
	;
	if v242 == int32(0) {
		v377 = v227
		v378 = v228
		v381 = v232
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+22)))
	v248 = v246 + v247
	v249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+74)))
	if v249 <= int32(0) {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+120)))
	if v253 < v249 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v263 = (v249 - int32(1)) & int32(_a_F_RelationBuildDesc_5)
	v264 = int32(100)
	base.MemoryCopy(m, v255+v256<<(uint(int32(3))%32)+v263*v264+int32(28), v248, v264)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	F_populate_compact_attribute(m, v271, v263)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+86)))
	if v274 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v277 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+16)) = uint8(v277)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+90)))
	if v279 == int32(115) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+17)) = uint8(v282)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+90)))
	v285 = v284
	goto L61
L60:
	;
	v285 = v279
	goto L61
L61:
	;
	if v285&int32(255) == int32(118) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+18)) = uint8(v290)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+87)))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+88)))
	if v293 != int32(1) {
		v369 = v232
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v372 = v228 + v292
	v374 = v227 - int32(1)
	if v374 != 0 {
		v227 = v374
		v228 = v372
		v232 = v369
		goto L49
	} else {
		goto L84
	}
L66:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v210)+52))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+18)))
	if base.Ui32(v298&int32(2047)) <= base.Ui32(int32(24)) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+287)))
	if v314 != 0 {
		v369 = v232
		goto L65
	} else {
		goto L73
	}
L68:
	;
	v306 = F_getmissingattr(m, v296, int32(25), v23+int32(287))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v311 = F_fastgetattr_3(m, v242, int32(25), v296, v23+int32(287))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L72
	}
L71:
	;
	v313 = v306
	goto L67
L72:
	;
	v313 = v311
	goto L67
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+288)) = int32(1)
	if v232 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v322 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321)+120)))
	v325 = F_MemoryContextAllocZero(m, v320, v322<<(uint(int32(4))%32))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	v327 = v232
	goto L76
L76:
	;
	v332 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+72)))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+82)))
	v334 = int32(*(*int8)(unsafe.Add(mBase, uint32(v248)+83)))
	v337 = F_array_get_element(m, v313, int32(1), v23+int32(288), int32(-1), v332, v333, v334, v23+int32(159))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L78
	}
L77:
	;
	v327 = v325
	goto L76
L78:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+82)))
	if v339 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v366 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v263<<(uint(int32(4))%32)))) = uint8(v366)
	v369 = v327
	goto L65
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v327+v263<<(uint(int32(4))%32))+8)) = v337
	goto L79
L81:
	;
	goto L82
L82:
	;
	v346 = int32(_a_F_RelationBuildDesc_1)
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3]))
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v350
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+82)))
	v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+72)))
	v357 = F_datumCopy(m, v337, v355, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v327+v263<<(uint(int32(4))%32))+8)) = v357
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v347
	goto L79
L84:
	;
	goto L50
L85:
	;
	F_relation_close(m, v210, int32(1))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v377 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+16)))
	if v390 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v946 = int32(0)
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
	if v946 < v955 {
		goto L208
	} else {
		goto L209
	}
L89:
	;
	F_pfree(m, v192)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L4
	} else {
		goto L206
	}
L90:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
	v403 = base.B2i32(base.Ui32(v401) < base.Ui32(int32(_a_F_RelationBuildDesc_6)))
	goto L95
L91:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+17)))
	if v391 != 0 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+18)))
	if v381|(v392|base.B2i32(int32(0) < v378)) != 0 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v398 = int32(*(*int16)(unsafe.Add(mBase, uint32(v397)+122)))
	if v398 <= int32(0) {
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v404)+24)) = v192
	if int32(0) < v378 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192)+8)) = v381
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v621 = int32(*(*int16)(unsafe.Add(mBase, uint32(v620)+122)))
	if v621 <= int32(0) {
		goto L145
	} else {
		goto L146
	}
L97:
	;
	v408 = int32(0)
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v413 = F_MemoryContextAllocZero(m, v410, v378<<(uint(int32(3))%32))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v597 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v192)+12)) = uint16(v597)
	goto L96
L100:
	;
	v416 = v23 + int32(288)
	v420 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v178))))
	F_ScanKeyInit(m, v416, int32(2), int32(3), int32(184), v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v425 = F_table_open(m, int32(2604), int32(1))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L103
	}
L102:
	;
	F_systable_endscan(m, v431)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L130
	}
L103:
	;
	v428 = int32(1)
	v431 = F_systable_beginscan(m, v425, int32(2656), v428, int32(0), v428, v416)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v433 = F_systable_getnext(m, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	if v433 == int32(0) {
		v541 = v408
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v440 = v433
	v442 = v408
	goto L107
L107:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+22)))
	v459 = v457 + v458
	if v378 <= v442 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v541 = v532
	goto L102
L109:
	;
	v463 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v425)+52))
	v487 = F_fastgetattr_3(m, v440, int32(4), v484, v23+int32(287))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L116
	}
L112:
	;
	if v463 == int32(0) {
		v541 = v442
		goto L102
	} else {
		goto L113
	}
L113:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v468 = int32(*(*int16)(unsafe.Add(mBase, uint32(v459)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v467 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_7), v23+int32(112))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_8), int32(_a_F_RelationBuildDesc_9))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v541 = v442
	goto L102
L116:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+287)))
	if v489 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v534 = F_systable_getnext(m, v431)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L128
	}
L118:
	;
	v494 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v515 = F_text_to_cstring(m, base.I32_wrap_i64(v487))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L125
	}
L121:
	;
	if v494 == int32(0) {
		v532 = v442
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v459)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v498 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_10), v23+int32(96))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_11), int32(_a_F_RelationBuildDesc_9))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v532 = v442
	goto L117
L125:
	;
	v519 = v413 + v442<<(uint(int32(3))%32)
	v520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v459)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v519))) = uint16(v520)
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v524 = F_MemoryContextStrdup(m, v523, v515)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+4)) = v524
	F_pfree(m, v515)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v532 = v442 + int32(1)
	goto L117
L128:
	;
	if v534 != 0 {
		v440 = v534
		v442 = v532
		goto L107
	} else {
		goto L129
	}
L129:
	;
	goto L108
L130:
	;
	F_relation_close(m, v425, int32(1))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	if v541 == v378 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if int32(2) <= v541 {
		goto L138
	} else {
		goto L139
	}
L133:
	;
	v564 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	if v564 == int32(0) {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v378 - v541
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v568 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_12), v23+int32(80))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_13), int32(_a_F_RelationBuildDesc_9))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	goto L132
L138:
	;
	F_pg_qsort(m, v413, v541, int32(8), int32(1804))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v592))) = v413
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+24))
	*(*uint16)(unsafe.Add(mBase, uint32(v595)+12)) = uint16(v541)
	goto L96
L141:
	;
	goto L140
L142:
	;
	v917 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v902)+122)))
	if v917 != 0 {
		goto L88
	} else {
		goto L205
	}
L143:
	;
	v858 = int32(0)
	v859 = int32(*(*int16)(unsafe.Add(mBase, uint32(v843)+120)))
	if v859 <= v858 {
		v902 = v843
		goto L142
	} else {
		goto L198
	}
L144:
	;
	v636 = v23 + int32(288)
	v640 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v178))))
	F_ScanKeyInit(m, v636, int32(9), int32(3), int32(184), v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L4
	} else {
		goto L151
	}
L145:
	;
	if base.Ui32(v401) < base.Ui32(int32(_a_F_RelationBuildDesc_6)) {
		v902 = v620
		goto L142
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v632 = F_MemoryContextAllocZero(m, v629, v621*int32(12))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L150
	}
L148:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+16)))
	if v624 != int32(1) {
		v843 = v620
		goto L143
	} else {
		goto L149
	}
L149:
	;
	v634 = int32(0)
	goto L144
L150:
	;
	v634 = v632
	goto L144
L151:
	;
	v643 = int32(0)
	v646 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L4
	} else {
		goto L153
	}
L152:
	;
	F_systable_endscan(m, v652)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L185
	}
L153:
	;
	v649 = int32(1)
	v652 = F_systable_beginscan(m, v646, int32(2665), v649, int32(0), v649, v636)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v654 = F_systable_getnext(m, v652)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	if v654 == int32(0) {
		v781 = v643
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v661 = v654
	v663 = v643
	goto L157
L157:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v661)+16))
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678)+22)))
	v680 = v678 + v679
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+72)))
	switch v681 - int32(99) {
	case 0:
		goto L160
	default:
		v772 = v663
		goto L159
	case 11:
		goto L161
	}
L158:
	;
	v781 = v772
	goto L152
L159:
	;
	v774 = F_systable_getnext(m, v652)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L183
	}
L160:
	;
	if v621 <= v663 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+76)))
	if v684 != 0 {
		v772 = v663
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v685 = F_extractNotNullColumn(m, v661)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v691 = int32(105)
	*(*uint8)(unsafe.Add(mBase, uint32(v687+v685<<(uint(int32(3))%32))+27)) = uint8(v691)
	v772 = v663
	goto L159
L164:
	;
	v696 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L4
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v646)+52))
	v718 = F_fastgetattr_3(m, v661, int32(28), v715, v23+int32(287))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L171
	}
L167:
	;
	if v696 == int32(0) {
		v781 = v663
		goto L152
	} else {
		goto L168
	}
L168:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v700 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_14), v23-int32(-64))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_15), int32(_a_F_RelationBuildDesc_16))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	v781 = v663
	goto L152
L171:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+287)))
	if v720 == int32(1) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v725 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L4
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v744 = F_text_to_cstring(m, base.I32_wrap_i64(v718))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L4
	} else {
		goto L179
	}
L175:
	;
	if v725 == int32(0) {
		v772 = v663
		goto L159
	} else {
		goto L176
	}
L176:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v729 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_17), v23+int32(48))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_18), int32(_a_F_RelationBuildDesc_16))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	v772 = v663
	goto L159
L179:
	;
	v748 = v634 + v663*int32(12)
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v748)+8)) = uint8(v749)
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v748)+9)) = uint8(v751)
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+106)))
	*(*uint8)(unsafe.Add(mBase, uint32(v748)+10)) = uint8(v753)
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v759 = F_MemoryContextStrdup(m, v756, v680+int32(4))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v748))) = v759
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v764 = F_MemoryContextStrdup(m, v763, v744)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v748)+4)) = v764
	F_pfree(m, v744)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v772 = v663 + int32(1)
	goto L159
L183:
	;
	if v774 != 0 {
		v661 = v774
		v663 = v772
		goto L157
	} else {
		goto L184
	}
L184:
	;
	goto L158
L185:
	;
	F_relation_close(m, v646, int32(1))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	if v781 == v621 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	if int32(2) <= v781 {
		goto L193
	} else {
		goto L194
	}
L188:
	;
	v804 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	if v804 == int32(0) {
		goto L187
	} else {
		goto L190
	}
L190:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v621 - v781
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v808 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_19), v23+int32(32))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_20), int32(_a_F_RelationBuildDesc_16))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	goto L187
L193:
	;
	F_pg_qsort(m, v634, v781, int32(12), int32(1805))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L4
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v634
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)+24))
	*(*uint16)(unsafe.Add(mBase, uint32(v835)+14)) = uint16(v781)
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	if base.Ui32(v401) < base.Ui32(int32(_a_F_RelationBuildDesc_6)) {
		v902 = v837
		goto L142
	} else {
		goto L197
	}
L196:
	;
	goto L195
L197:
	;
	v843 = v837
	goto L143
L198:
	;
	v865 = v858
	v867 = v843
	goto L199
L199:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	v885 = v882 + v865<<(uint(int32(3))%32)
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+35)))
	if v886 == int32(117) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v902 = v892
	goto L142
L201:
	;
	v889 = int32(118)
	*(*uint8)(unsafe.Add(mBase, uint32(v885)+35)) = uint8(v889)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v892 = v891
	goto L203
L202:
	;
	v892 = v867
	goto L203
L203:
	;
	v894 = v865 + int32(1)
	v895 = int32(*(*int16)(unsafe.Add(mBase, uint32(v892)+120)))
	if v894 < v895 {
		v865 = v894
		v867 = v892
		goto L199
	} else {
		goto L204
	}
L204:
	;
	goto L200
L205:
	;
	v918 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v192)+14)) = uint16(v918)
	goto L88
L206:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v95)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v922)+24)) = int32(0)
	goto L88
L207:
	;
	v1033 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+128)) = v1033
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+88)) = uint8(v1033)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+84)) = v1033
	v1039 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v95)+92)) = v1039
	*(*int64)(unsafe.Add(mBase, uint32(v95)+100)) = v1039
	*(*int64)(unsafe.Add(mBase, uint32(v95)+108)) = v1039
	*(*int64)(unsafe.Add(mBase, uint32(v95)+116)) = v1039
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+124)) = uint8(v1033)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049)+119)))
	switch v1050 - int32(73) {
	case 0, 32:
		goto L228
	default:
		goto L226
	case 10, 36, 41, 43:
		goto L227
	}
L208:
	;
	v959 = v945 + int32(28)
	v966 = v946
	v967 = v955
	v969 = v946
	goto L212
L209:
	;
	v1023 = v946
	v1030 = v955
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v945)+20)) = v1030
	*(*int32)(unsafe.Add(mBase, uint32(v945)+16)) = v1023
	goto L207
L211:
	;
	v1023 = v1017
	v1030 = v996
	goto L210
L212:
	;
	v975 = v959 + v955<<(uint(int32(3))%32) + v966*int32(100)
	v978 = v959 + v966<<(uint(int32(3))%32)
	if v955 != v967 {
		v996 = v967
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1017 = v955
	goto L211
L214:
	;
	v997 = int32(*(*int16)(unsafe.Add(mBase, uint32(v978)+2)))
	if v997 <= int32(0) {
		v1017 = v966
		goto L211
	} else {
		goto L222
	}
L215:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978)+7)))
	if v980 != int32(118) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v996 = v966
	goto L214
L217:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978)+4)))
	if v983 != int32(1) {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978)+6)))
	if v986&int32(6) != 0 {
		goto L216
	} else {
		goto L219
	}
L219:
	;
	v989 = int32(*(*int16)(unsafe.Add(mBase, uint32(v978)+2)))
	if v989 <= int32(0) {
		goto L216
	} else {
		goto L220
	}
L220:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975)+90)))
	if v992 != int32(118) {
		v996 = v955
		goto L214
	} else {
		goto L221
	}
L221:
	;
	goto L216
L222:
	;
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975)+90)))
	if v1000 == int32(118) {
		v1017 = v966
		goto L211
	} else {
		goto L223
	}
L223:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978)+5)))
	v1009 = (v969 + v1003 - int32(1)) & (int32(0) - v1003)
	if int32(_a_F_RelationBuildDesc_21) < v1009 {
		v1017 = v966
		goto L211
	} else {
		goto L224
	}
L224:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v978))) = uint16(v1009)
	v1015 = v966 + int32(1)
	if v1015 != v955 {
		v966 = v1015
		v967 = v996
		v969 = v1009 + v997
		goto L212
	} else {
		goto L225
	}
L225:
	;
	goto L213
L226:
	;
	F_RelationParseRelOptions(m, v95, v78)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L4
	} else {
		goto L231
	}
L227:
	;
	F_RelationInitTableAccessMethod(m, v95)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L4
	} else {
		goto L230
	}
L228:
	;
	F_RelationInitIndexAccessInfo(m, v95)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L4
	} else {
		goto L229
	}
L229:
	;
	goto L226
L230:
	;
	goto L226
L231:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059)+124)))
	if v1060 == int32(1) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068)+125)))
	if v1069 == int32(1) {
		goto L238
	} else {
		goto L239
	}
L233:
	;
	F_RelationBuildRuleLock(m, v95)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L4
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v95)+68)) = int64(0)
	v1068 = v1059
	goto L232
L236:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v1068 = v1065
	goto L232
L237:
	;
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1077)+127)))
	if v1078 == int32(1) {
		goto L243
	} else {
		goto L244
	}
L238:
	;
	F_RelationBuildTriggers(m, v95)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L4
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+76)) = int32(0)
	v1077 = v1068
	goto L237
L241:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v1077 = v1074
	goto L237
L242:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+60)) = v1085
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[10]))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1090)+117)))
	if v1091 != 0 {
		goto L248
	} else {
		goto L249
	}
L243:
	;
	F_RelationBuildRowSecurity(m, v95)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L4
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+80)) = int32(0)
	goto L242
L246:
	;
	goto L242
L247:
	;
	F_RelationInitPhysicalAddr(m, v95)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L4
	} else {
		goto L251
	}
L248:
	;
	v1092 = int32(0)
	goto L250
L249:
	;
	v1092 = v1089
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+64)) = v1092
	goto L247
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
	F_pfree(m, v78)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0]))
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101+v53)+4)))
	if v1103 != int32(1) {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	F_RelationDestroyRelation(m, v95, int32(0))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0]))
	v1112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1110+v53)+4)) = uint8(v1112)
	v1116 = F_ScanPgRelation(m, l0, int32(1), v1112)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	if v1116 != 0 {
		v78 = v1116
		goto L14
	} else {
		goto L256
	}
L256:
	;
	goto L15
L257:
	;
	v1198 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+26)) = uint8(v1198)
	v1202 = v95
	goto L8
L258:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[11]))
	v1158 = F_hash_search(m, v1154, v178, int32(1), v23+int32(160))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L4
	} else {
		goto L259
	}
L259:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+160)))
	if v1160 == int32(1) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+4)) = v95
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+16))
	if v1165 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	goto L262
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+4)) = v95
	goto L257
L263:
	;
	F_RelationDestroyRelation(m, v1163, int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L4
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[12]))
	if v1172 == int32(0) {
		goto L257
	} else {
		goto L267
	}
L266:
	;
	goto L257
L267:
	;
	v1177 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	if v1177 == int32(0) {
		goto L257
	} else {
		goto L269
	}
L269:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1181 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_22), v23+int32(16))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(1314), int32(_a_F_RelationBuildDesc_4))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	goto L257
L272:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v1229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v1229
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v1228 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_23), v23+int32(144))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(589), int32(_a_F_RelationBuildDesc_24))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v1248
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v377
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_25), v23+int32(128))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L4
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(671), int32(_a_F_RelationBuildDesc_24))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCacheInitFilePostInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePostInvalidate[0]))
	F_LWLockRelease(m, v2+int32(2048))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_RelationCacheInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	v1 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[0]))
	if v13 == int32(_a_F_RelationCacheInvalidate_0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_read_relmap_file(m, int32(_a_F_RelationCacheInvalidate_1), int32(_a_F_RelationCacheInvalidate_2), int32(0), int32(22))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[1]))
	if v23 == int32(_a_F_RelationCacheInvalidate_0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[2]))
	F_read_relmap_file(m, int32(_a_F_RelationCacheInvalidate_3), v28, int32(0), int32(22))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v34 = v10 + int32(12)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[3]))
	F_hash_seq_init(m, v34, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v39 = F_hash_seq_search(m, v34)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	F_list_free(m, v311)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L127
	}
L12:
	;
	F_list_free(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L94
	}
L13:
	;
	v231 = v128
	goto L12
L14:
	;
	if v39 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = v39
	v44 = v1
	v45 = v1
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_smgrreleaseall(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L92
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	if v49 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F_smgrreleaseall(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L61
	}
L20:
	;
	v131 = F_hash_seq_search(m, v10+int32(12))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L59
	}
L21:
	;
	v55 = int32(_a_F_RelationCacheInvalidate_4)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[4])) = v57 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v61 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	if v52 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v127 = v44
	v128 = v45
	goto L20
L25:
	;
	goto L24
L26:
	;
	v127 = v124
	v128 = v125
	goto L20
L27:
	;
	F_RelationClearRelation(m, v48)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	switch v67 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L32
	default:
		goto L31
	}
L30:
	;
	v124 = v44
	v125 = v45
	goto L26
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v107 != int32(2662) {
		goto L48
	} else {
		goto L49
	}
L32:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+88))
	if v70 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v71 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	v76 = v74 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v76
	if v76 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	F_RelationInitPhysicalAddr(m, v48)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L46
	}
L37:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	F_smgrclose(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L45
	}
L38:
	;
	v81 = v71 + int32(76)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5]))
	if v83 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v90
	v92 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+80)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v81
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6])) = v81
	goto L40
L42:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6]))
	v90 = v85
	goto L41
L43:
	;
	goto L44
L44:
	;
	v87 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5])) = v87
	v90 = v87
	goto L41
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = int32(0)
	goto L36
L46:
	;
	goto L31
L47:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+25)))
	if v116 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	if v107 != int32(1259) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v114 = F_lappend(m, v45, v48)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	v112 = F_lcons(m, v48, v45)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v124 = v44
	v125 = v112
	goto L26
L53:
	;
	v124 = v44
	v125 = v114
	goto L26
L54:
	;
	v119 = F_lcons(m, v48, v44)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v121 = F_lappend(m, v44, v48)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	v124 = v119
	v125 = v45
	goto L26
L58:
	;
	v124 = v121
	v125 = v45
	goto L26
L59:
	;
	if v131 != 0 {
		v41 = v131
		v44 = v127
		v45 = v128
		goto L18
	} else {
		goto L60
	}
L60:
	;
	goto L19
L61:
	;
	v135 = int32(0)
	if v128 == v135 {
		v231 = v135
		goto L12
	} else {
		goto L62
	}
L62:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v138 <= int32(0) {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	v144 = int32(0)
	goto L64
L64:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v150 = int32(2)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v144<<(uint(v150)%32))))
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[7]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	goto L68
L65:
	;
	goto L13
L66:
	;
	v209 = v144 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v209 < v210 {
		v144 = v209
		goto L64
	} else {
		goto L91
	}
L67:
	;
	F_RelationRebuildRelation(m, v153)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L90
	}
L68:
	;
	if v156 == v150 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+25)))
	if v159 != int32(1) {
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	if v165 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	if v162 != int32(1) {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+72))
	v170 = v168 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+72)) = v170
	if v170 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v153)+256))
	if v198 != 0 {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	F_smgrclose(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L85
	}
L78:
	;
	v175 = v165 + int32(76)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5]))
	if v177 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+76)) = v184
	v186 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+80)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v175
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6])) = v175
	goto L80
L82:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6]))
	v184 = v179
	goto L81
L83:
	;
	goto L84
L84:
	;
	v181 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5])) = v181
	v184 = v181
	goto L81
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = int32(0)
	goto L76
L86:
	;
	F_pfree(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+26)) = uint8(v201)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+256)) = v201
	goto L66
L89:
	;
	goto L88
L90:
	;
	goto L66
L91:
	;
	goto L65
L92:
	;
	F_list_free(m, int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v311 = v1
	goto L11
L94:
	;
	if v127 == int32(0) {
		v311 = v1
		goto L11
	} else {
		goto L95
	}
L95:
	;
	v236 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v237 <= v236 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v311 = v127
	goto L11
L97:
	;
	goto L98
L98:
	;
	v242 = v236
	goto L99
L99:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v248 = int32(2)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247+v242<<(uint(v248)%32))))
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[7]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	goto L103
L100:
	;
	v311 = v127
	goto L11
L101:
	;
	v307 = v242 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v307 < v308 {
		v242 = v307
		goto L99
	} else {
		goto L126
	}
L102:
	;
	F_RelationRebuildRelation(m, v251)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L125
	}
L103:
	;
	if v254 == v248 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+25)))
	if v257 != int32(1) {
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	if v263 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	if v260 != int32(1) {
		goto L102
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)+72))
	v268 = v266 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+72)) = v268
	if v268 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L111
L111:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v251)+256))
	if v296 != 0 {
		goto L121
	} else {
		goto L122
	}
L112:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	F_smgrclose(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L120
	}
L113:
	;
	v273 = v263 + int32(76)
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5]))
	if v275 != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	goto L112
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+76)) = v282
	v284 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+80)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v273
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6])) = v273
	goto L115
L117:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6]))
	v282 = v277
	goto L116
L118:
	;
	goto L119
L119:
	;
	v279 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5])) = v279
	v282 = v279
	goto L116
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+12)) = int32(0)
	goto L111
L121:
	;
	F_pfree(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v299 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v251)+26)) = uint8(v299)
	*(*int32)(unsafe.Add(mBase, uint32(v251)+256)) = v299
	goto L101
L124:
	;
	goto L123
L125:
	;
	goto L101
L126:
	;
	goto L100
L127:
	;
	v319 = int32(0)
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[8]))
	if v321 <= v319 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	m.G0 = v10 + int32(32)
	return
L129:
	;
	v325 = v321 & int32(7)
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[9]))
	if base.Ui32(int32(8)) <= base.Ui32(v321) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v333 = v319
	v336 = int32(0)
	goto L133
L131:
	;
	v366 = v319
	goto L132
L132:
	;
	v374 = v366
	v376 = int32(0)
	goto L137
L133:
	;
	v342 = v327 + v333<<(uint(int32(3))%32)
	v343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+60)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+52)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+44)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+36)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+28)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+20)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+12)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+4)) = uint8(v343)
	v359 = int32(8)
	v360 = v333 + v359
	v362 = v336 + v359
	if v362 != v321&int32(2147483640) {
		v333 = v360
		v336 = v362
		goto L133
	} else {
		goto L135
	}
L134:
	;
	if v325 == int32(0) {
		goto L128
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	v366 = v360
	goto L132
L137:
	;
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v374<<(uint(int32(3))%32))+4)) = uint8(v384)
	v389 = v376 + v384
	if v389 != v325 {
		v374 = v374 + v384
		v376 = v389
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L128
L139:
	;
	goto L138
}
func F_RelationClearMissing(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v6 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(288)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+120)))
	base.MemoryFill(m, v11+int32(80), int32(0), int32(200))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v6
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v33)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)) = uint8(v33)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v33)
	v41 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v57 = int64(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_relation_close(m, v41, int32(3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v62 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(v13), base.I64_extend16_s(v57))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v62 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v67)+88)))
	if v69 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	v79 = F_heap_modify_tuple(m, v62, v72, v11+int32(80), v11+int32(48), v11+int32(16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_ReleaseCatCache(m, v62)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	F_CatalogTupleUpdate(m, v41, v79+int32(4), v79)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v79)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v91 = v57 + int64(1)
	if v91 != base.I64_extend_i32_u(v15+int32(1))&int64(65535) {
		v57 = v91
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	m.G0 = v11 + int32(288)
	return
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v13
	*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v57)
	F_errmsg_internal(m, int32(_a_F_RelationClearMissing_0), v11)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_RelationClearMissing_1), int32(2017), int32(_a_F_RelationClearMissing_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationForgetRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[0]))
	v15 = F_hash_search(m, v10, v6+int32(12), v2, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 == int32(0) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v19 == int32(0) {
				m.G0 = v6 + int32(16)
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				if v22 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v84
						F_errmsg_internal(m, int32(_a_F_RelationForgetRelation_0), v6)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RelationForgetRelation_1), int32(2897), int32(_a_F_RelationForgetRelation_2))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
					if v23 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
						if v26 == int32(0) {
							F_RelationClearRelation(m, v19)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[1]))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v31
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							if v33 != 0 {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
								v38 = v36 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v38
								if v38 == int32(0) {
									v43 = v33 + int32(76)
									v45 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2]))
									if v45 != 0 {
										v47 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3]))
										v52 = v47
									} else {
										v49 = int32(_a_F_RelationForgetRelation_3)
										*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2])) = v49
										v52 = v49
									}
									*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v52
									v54 = int32(_a_F_RelationForgetRelation_3)
									*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v54
									*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v43
									*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3])) = v43
								} else {
								}
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
								F_smgrclose(m, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
									if v66 != 0 {
										F_pfree(m, v66)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
											*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
											m.G0 = v6 + int32(16)
											return
										}
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
								if v66 != 0 {
									F_pfree(m, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[1]))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v31
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if v33 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
							v38 = v36 - int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v38
							if v38 == int32(0) {
								v43 = v33 + int32(76)
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2]))
								if v45 != 0 {
									v47 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3]))
									v52 = v47
								} else {
									v49 = int32(_a_F_RelationForgetRelation_3)
									*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2])) = v49
									v52 = v49
								}
								*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v52
								v54 = int32(_a_F_RelationForgetRelation_3)
								*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v43
								*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3])) = v43
							} else {
							}
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							F_smgrclose(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
								if v66 != 0 {
									F_pfree(m, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
							if v66 != 0 {
								F_pfree(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v69 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_RelationGetFKeyList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v13 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v118
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v118 = v16
	goto L1
L3:
	;
	goto L4
L4:
	;
	v18 = v9 + int32(-56)
	v22 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v18, int32(9), int32(3), int32(184), v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v29 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = int32(1)
	v35 = F_systable_beginscan(m, v29, int32(2665), v32, int32(0), v32, v18)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = v2
	v43 = v37
	goto L13
L11:
	;
	v89 = v2
	goto L12
L12:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L23
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+22)))
	v49 = v47 + v48
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+72)))
	if v50 == int32(102) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v89 = v83
	goto L12
L15:
	;
	v54 = F_palloc0(m, int32(280))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	v83 = v42
	goto L17
L17:
	;
	v84 = F_systable_getnext(m, v35)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(478)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v62
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+20)) = uint8(v64)
	v74 = int32(0)
	F_DeconstructFkConstraintRow(m, v43, v54+int32(16), v54+int32(22), v54+int32(86), v54+int32(152), v74, v74, v74, v74)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v80 = F_lappend(m, v42, v54)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v83 = v80
	goto L17
L21:
	;
	if v84 != 0 {
		v42 = v83
		v43 = v84
		goto L13
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	F_relation_close(m, v29, int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v99 = int32(_a_F_RelationGetFKeyList_0)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[0]))
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[0])) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v106 = F_copyObjectImpl(m, v89)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v108)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v106
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[0])) = v100
	F_list_free_deep(m, v105)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v118 = v89
	goto L1
}
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v50 int32
	_ = v50
	v7 = F_palloc(m, int32(92))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		if v11 < l1 {
			v23 = F_palloc_mul(m, int32(56), l1)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v26 = v23
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v26
				if int32(0) < l2 {
					v31 = F_palloc_mul(m, int32(56), l2)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v31
						v34 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v34)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v34)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v33
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+69)))
						v42 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v42
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v41)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v42
						*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v42
						v50 = v41 ^ int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v50)
						return v7
					}
				} else {
					v33 = int32(0)
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v34)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v33
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+69)))
					v42 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v42
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v41)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v42
					*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v42
					v50 = v41 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v50)
					return v7
				}
			}
		} else {
			v26 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v26
			if int32(0) < l2 {
				v31 = F_palloc_mul(m, int32(56), l2)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = v31
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v34)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v33
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+69)))
					v42 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v42
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v41)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v42
					*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v42
					v50 = v41 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v50)
					return v7
				}
			} else {
				v33 = int32(0)
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v34)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v33
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+69)))
				v42 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v42
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v41)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v42
				*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v42
				v50 = v41 ^ int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v50)
				return v7
			}
		}
	}
}
func F_RelationGetSmgr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v14
		v16 = F_smgropen(m, v6, v11)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v16
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
			if v22 != 0 {
				v30 = v22
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v24
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
				v30 = v28
			}
			*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v30 + int32(1)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v35 = v34
			m.G0 = v6 + int32(16)
			return v35
		}
	} else {
		v35 = v8
		m.G0 = v6 + int32(16)
		return v35
	}
}
func F_RelationIncrementReferenceCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIncrementReferenceCount[0]))
	F_ResourceOwnerEnlarge(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6 + int32(1)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIncrementReferenceCount[1]))
		if v11 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIncrementReferenceCount[0]))
			F_ResourceOwnerRemember(m, v13, base.I64_extend_i32_u(l0), int32(_a_F_RelationIncrementReferenceCount_0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_RelationInitLockInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v2
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitLockInfo[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+117)))
	if v8 != 0 {
		v9 = int32(0)
	} else {
		v9 = v6
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v9
	return
}
func F_RelationIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v134
L2:
	;
	return int32(0)
L3:
	;
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v21)
	v134 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_RelationIsVisibleExt_0), v11)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_RelationIsVisibleExt_1), int32(941), int32(_a_F_RelationIsVisibleExt_2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v40 = v36 + v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	if v41 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L40
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIsVisibleExt[0]))
	v46 = int32(0)
	if v45 == v46 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIsVisibleExt[0]))
	if v88 == int32(0) {
		v124 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v84 == int32(0) {
		v124 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v84 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v52 <= int32(0) {
		v78 = v46
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v78
	goto L18
L23:
	;
	v55 = int32(0)
	if v55 < v52 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v58 = v52
	goto L26
L25:
	;
	v58 = v55
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v61 = int32(0)
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(int32(2))%32))))
	v70 = base.B2i32(v69 == v41)
	if v69 == v41 {
		v78 = v70
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v78 = v70
	goto L22
L29:
	;
	v72 = v61 + int32(1)
	if v72 != v58 {
		v61 = v72
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v92 <= v91 {
		v124 = v3
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v97 = v91
	goto L34
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v97<<(uint(int32(2))%32))))
	v110 = base.B2i32(v109 == v41)
	if v109 == v41 {
		v124 = v110
		goto L14
	} else {
		goto L36
	}
L35:
	;
	v124 = v110
	goto L14
L36:
	;
	v111 = F_get_relname_relid(m, v40+int32(4), v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v111 != 0 {
		v124 = v110
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v114 = v97 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v114 < v115 {
		v97 = v114
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v134 = v124
	goto L1
}
func F_RelationMapInvalidate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapInvalidate[0]))
		if v3 != int32(_a_F_RelationMapInvalidate_0) {
			return
		} else {
			F_read_relmap_file(m, int32(_a_F_RelationMapInvalidate_1), int32(_a_F_RelationMapInvalidate_2), int32(0), int32(22))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapInvalidate[1]))
		if v13 != int32(_a_F_RelationMapInvalidate_0) {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapInvalidate[2]))
			F_read_relmap_file(m, int32(_a_F_RelationMapInvalidate_3), v18, int32(0), int32(22))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_relation_needs_vacanalyze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 float32
	_ = v50
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 float64
	_ = v129
	var v131 int32
	_ = v131
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v137 float64
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 float64
	_ = v198
	var v200 int32
	_ = v200
	var v203 float64
	_ = v203
	var v205 float64
	_ = v205
	var v207 int32
	_ = v207
	var v208 float64
	_ = v208
	var v210 int32
	_ = v210
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v216 int32
	_ = v216
	var v218 float64
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 float64
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 float64
	_ = v243
	var v245 float64
	_ = v245
	var v248 int32
	_ = v248
	var v254 float64
	_ = v254
	var v256 float64
	_ = v256
	var v259 float64
	_ = v259
	var v260 float64
	_ = v260
	var v261 float64
	_ = v261
	var v263 int32
	_ = v263
	var v264 float64
	_ = v264
	var v266 float64
	_ = v266
	var v273 float64
	_ = v273
	var v275 float64
	_ = v275
	var v278 float64
	_ = v278
	var v279 float64
	_ = v279
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v287 float64
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int64
	_ = v304
	var v305 float32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v319 float32
	_ = v319
	var v321 float32
	_ = v321
	var v323 float32
	_ = v323
	var v326 float32
	_ = v326
	var v329 float32
	_ = v329
	var v332 float32
	_ = v332
	var v336 float32
	_ = v336
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v341 float64
	_ = v341
	var v342 float64
	_ = v342
	var v343 float32
	_ = v343
	var v346 float32
	_ = v346
	var v349 float64
	_ = v349
	var v351 float64
	_ = v351
	var v353 float64
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v369 float32
	_ = v369
	var v374 float32
	_ = v374
	var v378 float32
	_ = v378
	var v381 float32
	_ = v381
	var v385 float64
	_ = v385
	var v386 float64
	_ = v386
	var v388 float64
	_ = v388
	var v390 float64
	_ = v390
	var v396 int32
	_ = v396
	var v400 float32
	_ = v400
	var v401 float32
	_ = v401
	var v404 int32
	_ = v404
	var v408 float32
	_ = v408
	var v411 float32
	_ = v411
	var v415 float64
	_ = v415
	var v416 float64
	_ = v416
	var v418 float64
	_ = v418
	var v420 float64
	_ = v420
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 float64
	_ = v438
	var v439 float64
	_ = v439
	var v440 float64
	_ = v440
	var v441 float64
	_ = v441
	var v442 float64
	_ = v442
	var v466 int32
	_ = v466
	var v470 float64
	_ = v470
	var v471 float64
	_ = v471
	var v472 float64
	_ = v472
	var v473 float64
	_ = v473
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	v10 = int32(0)
	v40 = int64(0)
	v42 = m.G0
	v44 = v42 - int32(176)
	m.G0 = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+175)) = uint8(v10)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+108))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	v50 = *(*float32)(unsafe.Add(mBase, uint32(l2)+100))
	*(*int64)(unsafe.Add(mBase, uint32(l8)+40)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(l8)+32)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(l8)+24)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(l8)+16)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(l8)+8)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(l8))) = v40
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v10)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v10)
	if l1 != 0 {
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
		if v67 < l3 {
			v69 = v67
		} else {
			v69 = l3
		}
		if v67 < int32(0) {
			v72 = l3
		} else {
			v72 = v69
		}
		v74 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[0]))
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		if v75 < v74 {
			v77 = v75
		} else {
			v77 = v74
		}
		if v75 < int32(0) {
			v80 = v74
		} else {
			v80 = v77
		}
		v82 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[1]))
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v83 < int32(0) {
			v86 = v82
		} else {
			v86 = v83
		}
		v87 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
		v89 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[2]))
		if base.F64_ge(v87, float64(0)) != 0 {
			v92 = v87
		} else {
			v92 = v89
		}
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v95 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[3]))
		if int32(-2) < v93 {
			v98 = v93
		} else {
			v98 = v95
		}
		v99 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
		v101 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[4]))
		if base.F64_ge(v99, float64(0)) != 0 {
			v104 = v99
		} else {
			v104 = v101
		}
		v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v107 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[5]))
		if int32(-2) < v105 {
			v110 = v105
		} else {
			v110 = v107
		}
		v111 = *(*float64)(unsafe.Add(mBase, uint32(l1)+72))
		v113 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[6]))
		if base.F64_ge(v111, float64(0)) != 0 {
			v116 = v111
		} else {
			v116 = v113
		}
		v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v119 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[7]))
		v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v120 < int32(0) {
			v123 = v119
		} else {
			v123 = v120
		}
		v142 = v72
		v143 = v98
		v145 = v80
		v146 = v110
		v147 = v86
		v148 = v117
		v150 = v92
		v151 = v104
		v152 = v116
		v153 = v123
	} else {
		v125 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[0]))
		v127 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[1]))
		v129 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[2]))
		v131 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[3]))
		v133 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[4]))
		v135 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[5]))
		v137 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[6]))
		v140 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[7]))
		v142 = l3
		v143 = v131
		v145 = v125
		v146 = v135
		v147 = v127
		v148 = int32(1)
		v150 = v129
		v151 = v133
		v152 = v137
		v153 = v140
	}
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[8]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+140))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[9])))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[10])))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)+136))
	if base.Ui32(v161) < base.Ui32(int32(3)) {
		v181 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[11]))
		v182 = int32(0)
		if v156 == v182 {
			v192 = v181
			v193 = v182
		} else {
			if v181 == v142 {
				v188 = int32(1)
			} else {
				v188 = v142 - v181
			}
			v192 = v181
			v193 = int32(base.Ui32(v188+v156) >> (uint(int32(31)) % 32))
		}
	} else {
		v164 = v155 - v145
		v165 = int32(3)
		if base.Ui32(v164) < base.Ui32(v165) {
			v169 = v164 - v165
		} else {
			v169 = v164
		}
		if base.B2i32(base.Ui32(v169) < base.Ui32(int32(3)))|base.B2i32(int32(0) <= v161-v169) != 0 {
			v181 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[11]))
			v182 = int32(0)
			if v156 == v182 {
				v192 = v181
				v193 = v182
			} else {
				if v181 == v142 {
					v188 = int32(1)
				} else {
					v188 = v142 - v181
				}
				v192 = v181
				v193 = int32(base.Ui32(v188+v156) >> (uint(int32(31)) % 32))
			}
		} else {
			v177 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[11]))
			v192 = v177
			v193 = int32(1)
		}
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v193)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[12]))
	v198 = base.F64_convert_i32_s(v197)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[13]))
	v203 = base.F64_mul(base.F64_convert_i32_s(v200), float64(1.05))
	if base.F64_lt(v203, v198) != 0 {
		v205 = v198
	} else {
		v205 = v203
	}
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[14]))
	v208 = base.F64_convert_i32_s(v207)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[0]))
	v213 = base.F64_mul(base.F64_convert_i32_s(v210), float64(1.05))
	if base.F64_gt(v208, v213) != 0 {
		v215 = v208
	} else {
		v215 = v213
	}
	v216 = base.I32_trunc_sat_f64_s(v215)
	v218 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[15]))
	if base.F64_gt(v218, float64(1)) != 0 {
		v224 = base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i32_s(v216), v218))
	} else {
		v224 = v216
	}
	v225 = int32(1)
	if v142 <= v225 {
		v228 = v225
	} else {
		v228 = v142
	}
	v229 = base.I32_trunc_sat_f64_s(v205)
	v231 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[16]))
	if base.F64_gt(v231, float64(1)) != 0 {
		v237 = base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i32_s(v229), v231))
	} else {
		v237 = v229
	}
	if base.Ui32(int32(2)) < base.Ui32(v161) {
		v242 = v155 - v161
	} else {
		v242 = int32(0)
	}
	v243 = base.F64_convert_i32_u(v242)
	v245 = base.F64_div(v243, base.F64_convert_i32_s(v145))
	v248 = int32(0)
	if base.B2i32(base.F64_gt(v245, float64(1)) == v248)|base.B2i32(base.Ui32(v242) < base.Ui32(v224)) == v248 {
		v254 = float64(1)
		v256 = base.F64_div(v243, float64(1e+08))
		if base.F64_lt(v256, v254) != 0 {
			v259 = v254
		} else {
			v259 = v256
		}
		v260 = F_pow(m, v245, v259)
		mBase = m.M
		v261 = v260
	} else {
		v261 = v245
	}
	if v156 != 0 {
		v263 = v192 - v156
	} else {
		v263 = int32(0)
	}
	v264 = base.F64_convert_i32_u(v263)
	v266 = base.F64_div(v264, base.F64_convert_i32_u(v228))
	if base.B2i32(base.F64_gt(v266, float64(1)) == int32(0))|base.B2i32(base.Ui32(v263) < base.Ui32(v237)) != 0 {
		v281 = v266
	} else {
		v273 = float64(1)
		v275 = base.F64_div(v264, float64(1e+08))
		if base.F64_lt(v275, v273) != 0 {
			v278 = v273
		} else {
			v278 = v275
		}
		v279 = F_pow(m, v266, v278)
		mBase = m.M
		v281 = v279
	}
	v282 = base.F64_mul(v231, v281)
	*(*float64)(unsafe.Add(mBase, uint32(l8)+16)) = v282
	v284 = base.F64_mul(v218, v261)
	*(*float64)(unsafe.Add(mBase, uint32(l8)+8)) = v284
	if base.F64_lt(v282, v284) != 0 {
		v287 = v284
	} else {
		v287 = v282
	}
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = v287
	if v193 != 0 {
		v289 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
	} else {
	}
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[17]))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+117)))
	if v295 != 0 {
		v296 = int32(0)
	} else {
		v296 = v294
	}
	v300 = F_pgstat_fetch_entry(m, int32(2), v296, base.I64_extend_i32_u(l0), v44+int32(175))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		return
	} else {
		if v300 == int32(0) {
			m.G0 = v44 + int32(176)
			return
		} else {
			v304 = *(*int64)(unsafe.Add(mBase, uint32(v300)+80))
			v305 = float32(1)
			if base.Ui32(v48) < base.Ui32(v49) {
				v309 = v48
			} else {
				v309 = v49
			}
			v314 = int32(0)
			if base.B2i32(v49 <= v314)|base.B2i32(v48 <= v314) != 0 {
				v319 = v305
			} else {
				v319 = base.F32_sub(v305, base.F32_div(base.F32_convert_i32_u(v309), base.F32_convert_i32_u(v49)))
			}
			v321 = base.F32_convert_i64_s(v304)
			v323 = float32(0)
			if base.F32_lt(v50, v323) != 0 {
				v326 = v323
			} else {
				v326 = v50
			}
			v329 = base.F32_add(base.F32_mul(base.F32_demote_f64(v152), v326), base.F32_convert_i32_s(v153))
			if v146 < int32(0) {
				v336 = v329
			} else {
				v332 = base.F32_convert_i32_u(v146)
				if base.F32_gt(v329, v332) == int32(0) {
					v336 = v329
				} else {
					v336 = v332
				}
			}
			v338 = *(*int64)(unsafe.Add(mBase, uint32(v300)+88))
			v339 = *(*int64)(unsafe.Add(mBase, uint32(v300)+96))
			v341 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[18]))
			v342 = base.F64_promote_f32(v321)
			v343 = float32(1)
			if base.F32_gt(v336, v343) != 0 {
				v346 = v336
			} else {
				v346 = v343
			}
			v349 = base.F64_mul(v341, base.F64_div(v342, base.F64_promote_f32(v346)))
			*(*float64)(unsafe.Add(mBase, uint32(l8)+24)) = v349
			v351 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
			if base.F64_lt(v349, v351) != 0 {
				v353 = v351
			} else {
				v353 = v349
			}
			*(*float64)(unsafe.Add(mBase, uint32(l8))) = v353
			v356 = int32(0)
			v358 = v158 & v160 & v148
			if base.B2i32(base.F32_lt(v336, v321) == v356)|base.B2i32(v358 != int32(1)) == v356 {
				v364 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v364)
			} else {
			}
			v369 = base.F32_convert_i64_s(v339)
			v374 = base.F32_add(base.F32_mul(base.F32_mul(v326, base.F32_demote_f64(v151)), v319), base.F32_convert_i32_s(v143))
			if v143 < int32(0) {
			} else {
				v378 = float32(1)
				if base.F32_gt(v374, v378) != 0 {
					v381 = v374
				} else {
					v381 = v378
				}
				v385 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[19]))
				v386 = base.F64_mul(base.F64_div(base.F64_promote_f32(v369), base.F64_promote_f32(v381)), v385)
				*(*float64)(unsafe.Add(mBase, uint32(l8)+32)) = v386
				v388 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
				if base.F64_lt(v386, v388) != 0 {
					v390 = v388
				} else {
					v390 = v386
				}
				*(*float64)(unsafe.Add(mBase, uint32(l8))) = v390
				if v358&base.F32_lt(v374, v369) == int32(0) {
				} else {
					v396 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v396)
				}
			}
			v400 = base.F32_convert_i64_s(v338)
			v401 = base.F32_add(base.F32_mul(base.F32_demote_f64(v150), v326), base.F32_convert_i32_s(v147))
			if l0 == int32(2619) {
			} else {
				v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+119)))
				if v404 == int32(116) {
				} else {
					v408 = float32(1)
					if base.F32_gt(v401, v408) != 0 {
						v411 = v401
					} else {
						v411 = v408
					}
					v415 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[20]))
					v416 = base.F64_mul(base.F64_div(base.F64_promote_f32(v400), base.F64_promote_f32(v411)), v415)
					*(*float64)(unsafe.Add(mBase, uint32(l8)+40)) = v416
					v418 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
					if base.F64_lt(v416, v418) != 0 {
						v420 = v418
					} else {
						v420 = v416
					}
					*(*float64)(unsafe.Add(mBase, uint32(l8))) = v420
					if v358&base.F32_lt(v401, v400) == int32(0) {
					} else {
						v426 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v426)
					}
				}
			}
			v431 = F_errstart(m, l4, int32(0))
			mBase = m.M
			v432 = m.ExcPending
			if v432 != 0 {
				return
			} else {
				if int32(0) <= v143 {
					if v431 == int32(0) {
						v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+175)))
						if v506 != int32(1) {
							m.G0 = v44 + int32(176)
							return
						} else {
							F_pfree(m, v300)
							mBase = m.M
							v510 = m.ExcPending
							if v510 != 0 {
								return
							} else {
								m.G0 = v44 + int32(176)
								return
							}
						}
					} else {
						v438 = *(*float64)(unsafe.Add(mBase, uint32(l8)+24))
						v439 = *(*float64)(unsafe.Add(mBase, uint32(l8)+32))
						v440 = *(*float64)(unsafe.Add(mBase, uint32(l8)+40))
						v441 = *(*float64)(unsafe.Add(mBase, uint32(l8)+8))
						v442 = *(*float64)(unsafe.Add(mBase, uint32(l8)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v44)+88)) = v442
						*(*float64)(unsafe.Add(mBase, uint32(v44)+80)) = v441
						*(*float64)(unsafe.Add(mBase, uint32(v44)+72)) = v440
						*(*float64)(unsafe.Add(mBase, uint32(v44-int32(-64)))) = base.F64_promote_f32(v401)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+56)) = base.F64_promote_f32(v400)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+48)) = v439
						*(*float64)(unsafe.Add(mBase, uint32(v44)+40)) = base.F64_promote_f32(v374)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+32)) = base.F64_promote_f32(v369)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+24)) = v438
						*(*float64)(unsafe.Add(mBase, uint32(v44)+16)) = base.F64_promote_f32(v336)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+8)) = v342
						*(*int32)(unsafe.Add(mBase, uint32(v44))) = l2 + int32(4)
						F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_0), v44)
						mBase = m.M
						v466 = m.ExcPending
						if v466 != 0 {
							return
						} else {
							v498 = int32(3330)
							F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v498, int32(_a_F_relation_needs_vacanalyze_2))
							mBase = m.M
							v501 = m.ExcPending
							if v501 != 0 {
								return
							} else {
								v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+175)))
								if v506 != int32(1) {
									m.G0 = v44 + int32(176)
									return
								} else {
									F_pfree(m, v300)
									mBase = m.M
									v510 = m.ExcPending
									if v510 != 0 {
										return
									} else {
										m.G0 = v44 + int32(176)
										return
									}
								}
							}
						}
					}
				} else {
					if v431 == int32(0) {
						v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+175)))
						if v506 != int32(1) {
							m.G0 = v44 + int32(176)
							return
						} else {
							F_pfree(m, v300)
							mBase = m.M
							v510 = m.ExcPending
							if v510 != 0 {
								return
							} else {
								m.G0 = v44 + int32(176)
								return
							}
						}
					} else {
						v470 = *(*float64)(unsafe.Add(mBase, uint32(l8)+24))
						v471 = *(*float64)(unsafe.Add(mBase, uint32(l8)+40))
						v472 = *(*float64)(unsafe.Add(mBase, uint32(l8)+8))
						v473 = *(*float64)(unsafe.Add(mBase, uint32(l8)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v44)+160)) = v473
						*(*float64)(unsafe.Add(mBase, uint32(v44)+152)) = v472
						*(*float64)(unsafe.Add(mBase, uint32(v44)+144)) = v471
						*(*float64)(unsafe.Add(mBase, uint32(v44)+136)) = base.F64_promote_f32(v401)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+128)) = base.F64_promote_f32(v400)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+120)) = v470
						*(*float64)(unsafe.Add(mBase, uint32(v44)+112)) = base.F64_promote_f32(v336)
						*(*float64)(unsafe.Add(mBase, uint32(v44)+104)) = v342
						*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = l2 + int32(4)
						F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_3), v44+int32(96))
						mBase = m.M
						v492 = m.ExcPending
						if v492 != 0 {
							return
						} else {
							v498 = int32(3336)
							F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v498, int32(_a_F_relation_needs_vacanalyze_2))
							mBase = m.M
							v501 = m.ExcPending
							if v501 != 0 {
								return
							} else {
								v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+175)))
								if v506 != int32(1) {
									m.G0 = v44 + int32(176)
									return
								} else {
									F_pfree(m, v300)
									mBase = m.M
									v510 = m.ExcPending
									if v510 != 0 {
										return
									} else {
										m.G0 = v44 + int32(176)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		F_LockRelationOid(m, l0, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_RelationIdGetRelation(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
					v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+118)))
					if v15 == int32(116) {
						v18 = int32(_a_F_relation_open_0)
						v20 = *(*int32)(unsafe.Add(mBase, _c_F_relation_open[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_relation_open[0])) = v20 | int32(1)
					} else {
					}
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
					switch v26 - int32(83) {
					case 0, 22, 26, 29, 31, 33:
						v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_open[1])))
						if v30 == int32(0) {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+272))
							if v33 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = int32(0)
							} else {
							}
							v39 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
						} else {
							v36 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v36)
						}
					default:
						v39 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
					}
					m.G0 = v6 + int32(16)
					return v12
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
						F_errmsg_internal(m, int32(_a_F_relation_open_1), v6)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_relation_open_2), int32(62), int32(_a_F_relation_open_3))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
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
		v12 = F_RelationIdGetRelation(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+118)))
				if v15 == int32(116) {
					v18 = int32(_a_F_relation_open_0)
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_relation_open[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_relation_open[0])) = v20 | int32(1)
				} else {
				}
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
				switch v26 - int32(83) {
				case 0, 22, 26, 29, 31, 33:
					v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_open[1])))
					if v30 == int32(0) {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+272))
						if v33 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = int32(0)
						} else {
						}
						v39 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
					} else {
						v36 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v36)
					}
				default:
					v39 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
				}
				m.G0 = v6 + int32(16)
				return v12
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_relation_open_1), v6)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_relation_open_2), int32(62), int32(_a_F_relation_open_3))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
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
}
func F_relation_openrv(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(0)
			v15 = F_RangeVarGetRelidExtended(m, l0, l1, v12, v12, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_RelationIdGetRelation(m, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
						if v20 == int32(116) {
							v23 = int32(_a_F_relation_openrv_0)
							v25 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0])) = v25 | int32(1)
						} else {
						}
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
						switch v31 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv[1])))
							if v35 == int32(0) {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
								if v38 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = int32(0)
								} else {
								}
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
							} else {
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v41)
							}
						default:
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
						}
						m.G0 = v6 + int32(16)
						return v17
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
							F_errmsg_internal(m, int32(_a_F_relation_openrv_1), v6)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_relation_openrv_2), int32(62), int32(_a_F_relation_openrv_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
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
		}
	} else {
		v12 = int32(0)
		v15 = F_RangeVarGetRelidExtended(m, l0, l1, v12, v12, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_RelationIdGetRelation(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
					if v20 == int32(116) {
						v23 = int32(_a_F_relation_openrv_0)
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0])) = v25 | int32(1)
					} else {
					}
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
					switch v31 - int32(83) {
					case 0, 22, 26, 29, 31, 33:
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv[1])))
						if v35 == int32(0) {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
							if v38 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = int32(0)
							} else {
							}
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
						} else {
							v41 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v41)
						}
					default:
						v44 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
						*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
					}
					m.G0 = v6 + int32(16)
					return v17
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
						F_errmsg_internal(m, int32(_a_F_relation_openrv_1), v6)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_relation_openrv_2), int32(62), int32(_a_F_relation_openrv_3))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
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
	}
}
func F_relation_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(0)
			v16 = F_RangeVarGetRelidExtended(m, l0, l1, l2, v14, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16 != 0 {
					v18 = F_RelationIdGetRelation(m, v16)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v18 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
								F_errmsg_internal(m, int32(_a_F_relation_openrv_extended_0), v8)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_relation_openrv_extended_1), int32(62), int32(_a_F_relation_openrv_extended_2))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
							v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
							if v23 == int32(116) {
								v26 = int32(_a_F_relation_openrv_extended_3)
								v28 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0])) = v28 | int32(1)
							} else {
							}
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
							switch v34 - int32(83) {
							case 0, 22, 26, 29, 31, 33:
								v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv_extended[1])))
								if v38 == int32(0) {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
									if v41 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v41)+128)) = int32(0)
									} else {
									}
									v47 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
									*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
								} else {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v44)
								}
							default:
								v47 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
							}
							v52 = v18
							m.G0 = v8 + int32(16)
							return v52
						}
					}
				} else {
					v52 = int32(0)
					m.G0 = v8 + int32(16)
					return v52
				}
			}
		}
	} else {
		v14 = int32(0)
		v16 = F_RangeVarGetRelidExtended(m, l0, l1, l2, v14, v14)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 != 0 {
				v18 = F_RelationIdGetRelation(m, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v18 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
							F_errmsg_internal(m, int32(_a_F_relation_openrv_extended_0), v8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_relation_openrv_extended_1), int32(62), int32(_a_F_relation_openrv_extended_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
						if v23 == int32(116) {
							v26 = int32(_a_F_relation_openrv_extended_3)
							v28 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0])) = v28 | int32(1)
						} else {
						}
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
						switch v34 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv_extended[1])))
							if v38 == int32(0) {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
								if v41 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v41)+128)) = int32(0)
								} else {
								}
								v47 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v44)
							}
						default:
							v47 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
							*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
						}
						v52 = v18
						m.G0 = v8 + int32(16)
						return v52
					}
				}
			} else {
				v52 = int32(0)
				m.G0 = v8 + int32(16)
				return v52
			}
		}
	}
}
func F_validate_relation_as_table(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+119)))
	switch v8 - int32(83) {
	case 0, 19, 26, 29, 31, 33, 35:
		m.G0 = v5 + int32(16)
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v18 + int32(4)
				F_errmsg(m, int32(_a_F_validate_relation_as_table_0), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+119)))
					F_errdetail_relkind_not_supported(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_validate_relation_as_table_1), int32(152), int32(_a_F_validate_relation_as_table_2))
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
}
