package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CteScanNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+132))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	F_tuplestore_select_read_pointer(m, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v17*int32(24))+4)))
		v22 = int32(0)
		if base.B2i32(v21 == v22)|base.B2i32(v7 == int32(1)) == v22 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+136)))
			if v30 != 0 {
				v36 = int32(1)
				v39 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v36), v36, v15)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v39 != 0 {
						v78 = v15
						return v78
					} else {
						if v7 != int32(1) {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							m.T0[v74].(func(*base.Module, int32))(m, v15)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v78 = v15
								return v78
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
							if v45 != 0 {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
								m.T0[v74].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v78 = v15
									return v78
								}
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
								if v47 != 0 {
									F_ExecReScan(m, v46)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
										v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											if v51 != 0 {
												v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
												if v53&int32(2) == int32(0) {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v63)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v51)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
															m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int32(0)
															} else {
																v78 = v15
																return v78
															}
														}
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v59 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
													return int32(0)
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
									v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v51 != 0 {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
											if v53&int32(2) == int32(0) {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
												F_tuplestore_select_read_pointer(m, v9, v63)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_tuplestore_puttupleslot(m, v9, v51)
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
														m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int32(0)
														} else {
															v78 = v15
															return v78
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v59 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
											return int32(0)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v31 = int32(0)
				v33 = F_tuplestore_advance(m, v9, v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					if v33 != 0 {
						v36 = int32(1)
						v39 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v36), v36, v15)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 != 0 {
								v78 = v15
								return v78
							} else {
								if v7 != int32(1) {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
									m.T0[v74].(func(*base.Module, int32))(m, v15)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v78 = v15
										return v78
									}
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
									if v45 != 0 {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
										m.T0[v74].(func(*base.Module, int32))(m, v15)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v78 = v15
											return v78
										}
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
										if v47 != 0 {
											F_ExecReScan(m, v46)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
												v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return int32(0)
												} else {
													if v51 != 0 {
														v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
														if v53&int32(2) == int32(0) {
															v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
															F_tuplestore_select_read_pointer(m, v9, v63)
															mBase = m.M
															v65 = m.ExcPending
															if v65 != 0 {
																return int32(0)
															} else {
																F_tuplestore_puttupleslot(m, v9, v51)
																mBase = m.M
																v67 = m.ExcPending
																if v67 != 0 {
																	return int32(0)
																} else {
																	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
																	m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
																	mBase = m.M
																	v71 = m.ExcPending
																	if v71 != 0 {
																		return int32(0)
																	} else {
																		v78 = v15
																		return v78
																	}
																}
															}
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
															v59 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
															return int32(0)
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v59 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
														return int32(0)
													}
												}
											}
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
											v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												if v51 != 0 {
													v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
													if v53&int32(2) == int32(0) {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
														F_tuplestore_select_read_pointer(m, v9, v63)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															F_tuplestore_puttupleslot(m, v9, v51)
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
																v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
																m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
																mBase = m.M
																v71 = m.ExcPending
																if v71 != 0 {
																	return int32(0)
																} else {
																	v78 = v15
																	return v78
																}
															}
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
														v59 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
														return int32(0)
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v59 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v78 = v31
						return v78
					}
				}
			}
		} else {
			if v21 != 0 {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
				if v45 != 0 {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
					m.T0[v74].(func(*base.Module, int32))(m, v15)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v78 = v15
						return v78
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
					if v47 != 0 {
						F_ExecReScan(m, v46)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
							v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v51 != 0 {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
									if v53&int32(2) == int32(0) {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										F_tuplestore_select_read_pointer(m, v9, v63)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											F_tuplestore_puttupleslot(m, v9, v51)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
												m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													v78 = v15
													return v78
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v59 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
										return int32(0)
									}
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v59 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
									return int32(0)
								}
							}
						}
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
						v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 != 0 {
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
								if v53&int32(2) == int32(0) {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									F_tuplestore_select_read_pointer(m, v9, v63)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										F_tuplestore_puttupleslot(m, v9, v51)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
											m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v78 = v15
												return v78
											}
										}
									}
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v59 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
									return int32(0)
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v59 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
								return int32(0)
							}
						}
					}
				}
			} else {
				v36 = int32(1)
				v39 = F_tuplestore_gettupleslot(m, v9, base.B2i32(v7 == v36), v36, v15)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					if v39 != 0 {
						v78 = v15
						return v78
					} else {
						if v7 != int32(1) {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							m.T0[v74].(func(*base.Module, int32))(m, v15)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v78 = v15
								return v78
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+136)))
							if v45 != 0 {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
								m.T0[v74].(func(*base.Module, int32))(m, v15)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v78 = v15
									return v78
								}
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
								if v47 != 0 {
									F_ExecReScan(m, v46)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
										v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											if v51 != 0 {
												v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
												if v53&int32(2) == int32(0) {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
													F_tuplestore_select_read_pointer(m, v9, v63)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_tuplestore_puttupleslot(m, v9, v51)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
															m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int32(0)
															} else {
																v78 = v15
																return v78
															}
														}
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v59 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
													return int32(0)
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
									v51 = m.T0[v50].(func(*base.Module, int32) int32)(m, v46)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v51 != 0 {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
											if v53&int32(2) == int32(0) {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
												F_tuplestore_select_read_pointer(m, v9, v63)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_tuplestore_puttupleslot(m, v9, v51)
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
														m.T0[v69].(func(*base.Module, int32, int32))(m, v15, v51)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return int32(0)
														} else {
															v78 = v15
															return v78
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v59 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
												return int32(0)
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v59 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v58)+136)) = uint8(v59)
											return int32(0)
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
}
